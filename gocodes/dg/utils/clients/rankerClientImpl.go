package clients

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"codes/gocodes/dg/utils/log"

	"github.com/emirpasic/gods/maps/treebidimap"
	godsutils "github.com/emirpasic/gods/utils"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	globalmodels "codes/gocodes/dg/dg.model"
)

type RankerClientImpl struct {
	*GRPCClient
}

// type RankerBlackParam struct {
// 	dig.In
// 	RankerClient RankerClient `name:"RankerBlack"`
// }

// func NewRankerBlackClient(param RankerBlackParam) RankerClient {
// 	return param.RankerClient
// }

func NewRankerClient(addr string, timeout time.Duration) RankerClient {
	return &RankerClientImpl{
		NewGRPCClient(addr, timeout),
	}
}

func (rci *RankerClientImpl) AddRepo(repoID string, repoParams *RankerRepoParams) error {
	client, err := rci.ensureClient()
	if err != nil {
		return errors.WithStack(err)
	}
	rankerContext, err := newContext()
	if err != nil {
		return errors.WithStack(err)
	}
	req := &globalmodels.RankRepoOpRequest{
		Context: rankerContext,
		Repo: &globalmodels.RankRepoOperation{
			// -库ID, 必需. 增加库时需要保证唯一,否则返回错误. 其他操作需保证库ID存在, 否则返回错误
			RepoId: repoID,
			// -具体操作内容. 包括增加, 删除, 修改和查找
			Operation: globalmodels.RepoOperation_REPO_OPERATION_ADD,
			// -库级别. 在增加库和修改时必需, 其他操作不需要
			Level: globalmodels.RepoLevel(repoParams.RepoLevel),
			// -特征长度. 在增加库时必需且不可更新. 其他操作不需要
			FeatureLen: repoParams.FeatureLen,
			// -特征的数据类型, 目前包括float和short两种. 在增加库时必须且不可更新. 其他操作不需要.
			FeatureDataType: globalmodels.FeatureDataType(repoParams.FeatureDataType),
			// -库容量. 如果库级别定义为REPO_LEVEL_ON_GPU, 在增加库时必需, 表明库的最大容量. 其他级别和操作时不需要.
			Capacity: repoParams.RepoCapacity,
			// // -库当前大小, 仅在查询时作为返回值使用
			// Size int32 `protobuf:"varint,7,opt,name=Size,json=size" json:"Size"`
			// // -可选参数. 目前可用参数包括:
			// // -DynamicLoadNumber, 合法的数字型. 在库级别定义为REPO_LEVEL_ON_GPU时有效. 表明在启动时,按照数据新旧,最多加载到显存中的数据量
			// // -GPUThreads, [1,1,0,1]字符串格式. 在库级别定义为REPO_LEVEL_ON_GPU时必需. 表明数据在多个GPU上的分布. [1,1,0,1]表示当前服务器有三个GPU卡, 但是数据被平均存储在0,1和3号GPU卡上.
			Params: map[string]string{"GPUThreads": repoParams.GPU},
		},
	}
	start := time.Now()
	log.Debugf("Sending request to rpc://%v/AddRepo, req = %v", rci.GetAddr(), req)
	resp, err := client.RepoOperation(rci.GetContext(), req)
	if err != nil {
		log.Errorf("Failed to send request to rpc://%v/AddRepo\n%+v", rci.GetAddr(), err)
		return errors.WithStack(err)
	}
	log.Debugf("Got response from rpc://%v/AddRepo, resp = %v, took: %v", rci.GetAddr(), resp, time.Since(start))
	return rci.validateResponse(resp)
}

func (rci *RankerClientImpl) BatchRankFeature(queries []*RankFeatureQuery) (*treebidimap.Map, error) {
	return rci.batchRankFeature(rci.GetContext(), queries)
}

func (rci *RankerClientImpl) batchRankFeature(ctx context.Context, queries []*RankFeatureQuery) (*treebidimap.Map, error) {
	results := make([][]*globalmodels.RankItem, len(queries))
	g, _ := errgroup.WithContext(ctx)
	for i, query := range queries {
		i, query := i, query
		g.Go(func() error {
			items, err := rci.rankFeature(ctx, query)
			if err != nil {
				return errors.WithStack(err)
			}
			results[i] = items
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, errors.WithStack(err)
	}

	// sort value by score
	m := treebidimap.NewWith(godsutils.StringComparator, func(a, b interface{}) int {
		s1 := a.(*globalmodels.RankItem).Score
		s2 := b.(*globalmodels.RankItem).Score
		if s1 > s2 {
			return -1
		} else if s1 < s2 {
			return 1
		} else {
			return godsutils.StringComparator(a.(*globalmodels.RankItem).Id, b.(*globalmodels.RankItem).Id)
		}
	})
	for _, r := range results {
		for _, item := range r {
			if val, exist := m.Get(item.Id); exist {
				if item.Score > val.(*globalmodels.RankItem).Score {
					m.Put(item.Id, item)
				}
			} else {
				m.Put(item.Id, item)
			}
		}
	}
	return m, nil
}

func (rci *RankerClientImpl) BatchAddRankerFeature(params []*RankerFace, repoID, location string) error {
	return nil
}

func (rci *RankerClientImpl) BatchDeleteRankerFeature(repoID string, imageIDs []string) error {
	return nil
}

func (rci *RankerClientImpl) DeleteRepo(repoID string) error {
	return nil
}

func (rci *RankerClientImpl) QueryRepo(repoID string) (bool, error) {
	return false, nil
}

func (rci *RankerClientImpl) RankFeature(query *RankFeatureQuery) ([]*globalmodels.RankItem, error) {
	return rci.rankFeature(rci.GetContext(), query)
}

func (rci *RankerClientImpl) rankFeature(ctx context.Context, query *RankFeatureQuery) ([]*globalmodels.RankItem, error) {
	client, err := rci.ensureClient()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rankerContext, err := newContext()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	params := make(map[string]string, 0)
	if query.Confidence > 0 {
		params["ScoreThreshold"] = strconv.FormatFloat(float64(query.Confidence), 'f', -1, 32)
	}
	if query.TopX > 0 {
		max := strconv.Itoa(query.TopX)
		params["MaxCandidates"] = max
		params["PageSize"] = max // ranker的默认PageSize改成了100_(:з」∠)_
	}
	if query.SensorIds != nil {
		params["Locations"] = strings.Join(query.SensorIds, ",")
	} else {
		params["Locations"] = "0"
	}
	params["Normalization"] = "true"
	params["repoID"] = query.RepoId
	if query.StartTimestamp > 0 {
		params["StartTime"] = fmt.Sprintf("%v", query.StartTimestamp)
	} else {
		params["StartTime"] = "0"
	}
	if query.EndTimestamp > 0 {
		params["EndTime"] = fmt.Sprintf("%v", query.EndTimestamp)
	} else {
		params["EndTime"] = "9999999999999"
	}
	req := &globalmodels.RankFeatureRequest{
		Context: rankerContext,
		ObjectFeature: &globalmodels.ObjectProperty{
			Feature: query.Feature,
		},
		// -比对参数列表, 可选值如下
		// -ScoreThreshold, 数据类型为float, 指定比对分数的最小阈值, 小于该值不返回, 默认为0表示不过滤
		// -MaxCandidates 数据类型为int, 指定返回Top K
		// -PageSize 数据类型为int, 分页返回值, 指定每页大小. 默认为0表示不分页
		// -PageIndex 数据类型为int, 分页页数
		// -Normalization. 数据类型为bool, 指定是否需要对分数计算结果进行归一化处理. 默认为false
		// -ShowAttributes. 数据类型为bool, 指定是否返回比对结果的详细属性,比如时间,地点等信息. 默认为false
		// -FilterABC. 数据类型为int, 并可使用逗号分割传入多个值表示一个集合. 用于动态属性过滤, FliterABC表示过滤自定义的ABC属性, ABC属性值在传入的集合中
		// -RangeXYZ. 数据类型为int-int, 表示一个值的范围. 用于动态属性过滤, FliterXYZ表示过滤自定义的XYZ属性, 属性值在传入的范围之间.
		// -repoID 必需. 数据类型为int, 指定在哪个库中进行比对.
		// -Locations 必需. 数据类型为int, 并可使用逗号分割传入多个值. 比如"[1,2,3]"表示只在地点是1,2,3的特征中进行比对
		// -StartTime 必需. 数据类型为int64.// -StartTime 必需. 数据类型为int64.
		// -EndTime 必需. 数据类型为int64, 与StartTime配合使用,指定被特征的时间范围.
		Params: params,
	}
	start := time.Now()
	log.Debugf("Sending request to rpc://%v/RankFeature, req = %v", rci.GetAddr(), req)
	resp, err := client.RankFeature(ctx, req)
	if err != nil {
		log.Errorf("Failed to send request to rpc://%v/RankFeature\n%+v", rci.GetAddr(), err)
		return nil, errors.WithStack(err)
	}
	log.Debugf("Got response from rpc://%v/RankFeature, resp = %v, took: %v", rci.GetAddr(), resp, time.Since(start))
	err = rci.validateResponse(resp)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return resp.Candidates, nil
}

func (rci *RankerClientImpl) QueryFeatureEx(repoID string, id string) (*globalmodels.ObjectProperty, error) {
	resp, err := rci.QueryFeature(repoID, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if resp == nil || len(resp.ObjectFeatures) != 1 {
		return nil, errors.New("Invalid query feature response")
	}
	return resp.ObjectFeatures[0], nil
}

func (rci *RankerClientImpl) QueryFeature(repoID string, id string) (*globalmodels.RankFeatureOperation, error) {
	client, err := rci.ensureClient()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	rankerContext, err := newContext()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req := &globalmodels.RankFeatureOpRequest{
		Context: rankerContext,
		Features: &globalmodels.RankFeatureOperation{
			Operation: globalmodels.ObjectOperation_OBJECT_OPERATION_QUERY,
			RepoId:    repoID,
			ObjectFeatures: []*globalmodels.ObjectProperty{
				&globalmodels.ObjectProperty{
					Id: id,
				},
			},
		},
	}
	start := time.Now()
	log.Debugf("Sending request to rpc://%v/QueryFeature, req = %v", rci.GetAddr(), req)
	resp, err := client.ObjectFeatureOperation(rci.GetContext(), req)
	if err != nil {
		log.Errorf("Failed to send request to rpc://%v/QueryFeature\n%+v", rci.GetAddr(), err)
		return nil, errors.WithStack(err)
	}
	log.Debugf("Got response from rpc://%v/QueryFeature, resp = %v, took: %v", rci.GetAddr(), resp, time.Since(start))
	if err := rci.validateResponse(resp); err != nil {
		return nil, errors.WithStack(err)
	}
	return resp.Features, nil
}

func (rci *RankerClientImpl) ensureClient() (globalmodels.SimilarityServiceClient, error) {
	conn, err := rci.EnsureConn()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return globalmodels.NewSimilarityServiceClient(conn), nil
}

func (rci *RankerClientImpl) validateResponse(resp RankerResponse) error {
	// context := resp.GetContext()
	// if context == nil {
	// 	return errors.New("Nil response")
	// }
	// todo status check
	return nil
}

func newContext() (*globalmodels.RankRequestContext, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &globalmodels.RankRequestContext{
		SessionId: id.String(),
	}, nil
}
