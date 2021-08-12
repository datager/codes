package clients

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"

	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/log"

	"go.uber.org/dig"
	"golang.org/x/sync/errgroup"
	dg_model "codes/gocodes/dg/dg.model"

	"github.com/emirpasic/gods/maps/treebidimap"
	godsutils "github.com/emirpasic/gods/utils"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"google.golang.org/grpc"
)

const (
	RepoTypeBlack = iota + 1
	RepoTypeRuntime
)

const (
	FseReponseStatusOK = "200"
)

type fseConnection struct {
	client dg_model.SimilarityServiceClient
	conn   *grpc.ClientConn
}

type FsePoolService struct {
	addr string

	connectNubmer int

	timeoutSecond time.Duration
	tryTimes      int

	connectionPool []*fseConnection

	serviceOperationLock sync.RWMutex
	connectIndex         int
	connectLock          sync.Mutex
	isServiceOpen        bool
}

type RankerBlackParam struct {
	dig.In
	RankerClient RankerClient `name:"FseBlack"`
}

func NewRankerBlackClient(param RankerBlackParam) RankerClient {
	return param.RankerClient
}

func NewFsePoolService(addr string, connectNumber int, timeoutSecond time.Duration, tryTimes int) *FsePoolService {
	if addr == "" {
		panic("addr is empty")
	}

	service := &FsePoolService{
		addr:          addr,
		connectNubmer: connectNumber,
		timeoutSecond: timeoutSecond,
		tryTimes:      tryTimes,
	}

	service.init()

	return service
}

func (fps *FsePoolService) init() {
	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	if fps.isServiceOpen {
		return
	}

	for {

		if len(fps.connectionPool) < fps.connectNubmer {
			conn, err := grpc.Dial(fps.addr, grpc.WithInsecure())
			if err != nil {
				log.Errorln(" fse connect err:", err)
				<-time.After(time.Second)
				continue
			}

			connClient := dg_model.NewSimilarityServiceClient(conn)
			fps.connectionPool = append(fps.connectionPool, &fseConnection{
				conn:   conn,
				client: connClient,
			})

			// log.Infoln("create a new connect to fse black, total connect num is ", len(fps.blackConnectionPool))
			continue
		}

		break
	}

	fps.isServiceOpen = true
}

func (fps *FsePoolService) Close() error {
	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	if !fps.isServiceOpen {
		return nil
	}

	for _, connection := range fps.connectionPool {
		connection.client = nil
		connection.conn.Close()
	}

	fps.connectionPool = []*fseConnection{}

	fps.isServiceOpen = false

	return nil
}

func (fps *FsePoolService) AddRepo(repoID string, repoParams *RankerRepoParams) error {
	return fps.add(repoID, repoParams)
}

func (fps *FsePoolService) DeleteRepo(repoID string) error {
	if repoID == "" {
		return errors.New("invalid argument")
	}

	start := time.Now()

	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	conn := fps.connectionPool[fps.getConnectIndex()]
	rankerContext, err := fps.newContext()
	if err != nil {
		return errors.WithStack(err)
	}

	req := &dg_model.RankRepoOpRequest{
		Context: rankerContext,
		Repo: &dg_model.RankRepoOperation{
			RepoId:    repoID,
			Operation: dg_model.RepoOperation_REPO_OPERATION_DELETE,
		},
	}

	for {
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/DeleteRepo, req = %v", fps.addr, req)

		resp, err := conn.client.RepoOperation(ctx, req)
		if err != nil || resp.Context.Status != FseReponseStatusOK {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/DeleteRepo\n%v", fps.addr, err)
			return errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/DeleteRepo, resp = %v, took: %v", fps.addr, resp, time.Since(start))

		break
	}

	return nil
}

func (fps *FsePoolService) RankFeature(query *RankFeatureQuery) ([]*dg_model.RankItem, error) {
	return fps.rankFeature(query)
}

func (fps *FsePoolService) BatchRankFeature(queries []*RankFeatureQuery) (*treebidimap.Map, error) {
	ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
	defer cancel()

	return fps.batchRankFeature(ctx, queries)
}

func (fps *FsePoolService) QueryFeature(repoID string, id string) (*dg_model.RankFeatureOperation, error) {
	rankerContext, err := fps.newContext()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req := &dg_model.RankFeatureOpRequest{
		Context: rankerContext,
		Features: &dg_model.RankFeatureOperation{
			Operation: dg_model.ObjectOperation_OBJECT_OPERATION_QUERY,
			RepoId:    repoID,
			ObjectFeatures: []*dg_model.ObjectProperty{
				&dg_model.ObjectProperty{
					Id: id,
				},
			},
		},
	}

	start := time.Now()
	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	var resp *dg_model.RankFeatureOpResponse
	conn := fps.connectionPool[fps.getConnectIndex()]

	for {
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/QueryFeature, req = %v", fps.addr, req)

		resp, err = conn.client.ObjectFeatureOperation(ctx, req)
		if err != nil || resp.Context.Status != FseReponseStatusOK {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/QueryFeature\n%v", fps.addr, err)
			return nil, errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/QueryFeature, resp = %v, took: %v", fps.addr, resp, time.Since(start))

		break
	}

	return resp.Features, nil
}

func (fps *FsePoolService) QueryFeatureEx(repoID string, id string) (*dg_model.ObjectProperty, error) {
	resp, err := fps.QueryFeature(repoID, id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if resp == nil || len(resp.ObjectFeatures) != 1 {
		return nil, errors.New("Invalid query feature response")
	}
	return resp.ObjectFeatures[0], nil
}

func (fps *FsePoolService) BatchDeleteRankerFeature(repoID string, imageIDs []string) error {
	if repoID == "" || len(imageIDs) == 0 {
		return errors.New("invalid argument")
	}

	start := time.Now()

	objectProperty := make([]*dg_model.ObjectProperty, 0)
	for _, v := range imageIDs {
		obj := &dg_model.ObjectProperty{
			Id: v,
		}

		objectProperty = append(objectProperty, obj)
	}

	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	conn := fps.connectionPool[fps.getConnectIndex()]
	rankerContext, err := fps.newContext()
	if err != nil {
		return errors.WithStack(err)
	}

	req := &dg_model.RankFeatureOpRequest{
		Context: rankerContext,

		Features: &dg_model.RankFeatureOperation{
			RepoId:         repoID,
			Operation:      dg_model.ObjectOperation_OBJECT_OPERATION_DELETE,
			ObjectFeatures: objectProperty,
		},
	}

	for {
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/BatchDeleteRankerFeature, req = %v", fps.addr, req)

		resp, err := conn.client.ObjectFeatureOperation(ctx, req)
		if err != nil || resp.Context.Status != FseReponseStatusOK {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/BatchDeleteRankerFeature\n%v", fps.addr, err)
			return errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/BatchDeleteRankerFeature, resp = %v, took: %v", fps.addr, resp, time.Since(start))

		break
	}

	return nil
}

func (fps *FsePoolService) BatchAddRankerFeature(params []*RankerFace, repoID, location string) error {
	if params == nil {
		return errors.New("add ranker face params is nil")
	}

	start := time.Now()

	objectProperty := make([]*dg_model.ObjectProperty, 0)
	for _, v := range params {
		obj := &dg_model.ObjectProperty{
			Id:       v.ImageID,
			Location: location,
			Time:     utils.GetNowTs(),
			Feature:  v.Feature,
		}

		objectProperty = append(objectProperty, obj)
	}

	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	conn := fps.connectionPool[fps.getConnectIndex()]

	rankerContext, err := fps.newContext()
	if err != nil {
		return errors.WithStack(err)
	}

	req := &dg_model.RankFeatureOpRequest{
		Context: rankerContext,
		Features: &dg_model.RankFeatureOperation{
			RepoId:         repoID,
			Operation:      dg_model.ObjectOperation_OBJECT_OPERATION_ADD,
			ObjectFeatures: objectProperty,
		},
	}

	for {
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/AddRankerFeature, req = %v", fps.addr, req)

		resp, err := conn.client.ObjectFeatureOperation(ctx, req)
		if err != nil || resp.Context.Status != FseReponseStatusOK {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/AddRankerFeature\n%v", fps.addr, err)
			return errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/RankFeature, resp = %v, took: %v", fps.addr, resp, time.Since(start))

		break
	}

	return nil
}

func (fps *FsePoolService) QueryRepo(repoID string) (bool, error) {
	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	conn := fps.connectionPool[fps.getConnectIndex()]
	addr := fps.addr

	rankerContext, err := fps.newContext()
	if err != nil {
		return false, errors.WithStack(err)
	}

	req := &dg_model.RankRepoOpRequest{
		Context: rankerContext,
		Repo: &dg_model.RankRepoOperation{
			RepoId:    repoID,
			Operation: dg_model.RepoOperation_REPO_OPERATION_QUERY,
		},
	}

	start := time.Now()

	var exist bool

	for {
		log.Debugf("Sending request to rpc://%v//QueryRepo, req = %v", addr, req)
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		defer cancel()

		resp, err := conn.client.RepoOperation(ctx, req)
		if err != nil {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/QueryRepo\n%v", addr, err)
			return false, errors.WithStack(err)
		}

		if resp.Context.Status == FseReponseStatusOK {
			exist = true
		}

		log.Debugf("Got response from rpc://%v/QueryRepo, resp = %v, took: %v", addr, resp, time.Since(start))

		break
	}

	return exist, nil
}

func (fps *FsePoolService) batchRankFeature(ctx context.Context, queries []*RankFeatureQuery) (*treebidimap.Map, error) {
	results := make([][]*dg_model.RankItem, len(queries))
	g, _ := errgroup.WithContext(ctx)
	for i, query := range queries {
		i, query := i, query
		g.Go(func() error {
			items, err := fps.rankFeatureWithContext(ctx, query)
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
		s1 := a.(*dg_model.RankItem).Score
		s2 := b.(*dg_model.RankItem).Score
		if s1 > s2 {
			return -1
		} else if s1 < s2 {
			return 1
		} else {
			return godsutils.StringComparator(a.(*dg_model.RankItem).Id, b.(*dg_model.RankItem).Id)
		}
	})
	for _, r := range results {
		for _, item := range r {
			if val, exist := m.Get(item.Id); exist {
				if item.Score > val.(*dg_model.RankItem).Score {
					m.Put(item.Id, item)
				}
			} else {
				m.Put(item.Id, item)
			}
		}
	}
	return m, nil
}

func (fps *FsePoolService) add(repoID string, repoParams *RankerRepoParams) error {
	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	var conn *fseConnection
	var addr string

	conn = fps.connectionPool[fps.getConnectIndex()]
	addr = fps.addr

	rankerContext, err := fps.newContext()
	if err != nil {
		return errors.WithStack(err)
	}

	req := &dg_model.RankRepoOpRequest{
		Context: rankerContext,
		Repo: &dg_model.RankRepoOperation{
			// -库ID, 必需. 增加库时需要保证唯一,否则返回错误. 其他操作需保证库ID存在, 否则返回错误
			RepoId: repoID,
			// -具体操作内容. 包括增加, 删除, 修改和查找
			Operation: dg_model.RepoOperation_REPO_OPERATION_ADD,
			// -库级别. 在增加库和修改时必需, 其他操作不需要
			Level: dg_model.RepoLevel(repoParams.RepoLevel),
			// -特征长度. 在增加库时必需且不可更新. 其他操作不需要
			FeatureLen: repoParams.FeatureLen,
			// -特征的数据类型, 目前包括float和short两种. 在增加库时必须且不可更新. 其他操作不需要.
			FeatureDataType: dg_model.FeatureDataType(repoParams.FeatureDataType),
			// -库容量. 如果库级别定义为REPO_LEVEL_ON_GPU, 在增加库时必需, 表明库的最大容量. 其他级别和操作时不需要.
			Capacity: repoParams.RepoCapacity,
			// // -库当前大小, 仅在查询时作为返回值使用
			// Size int32 `protobuf:"varint,7,opt,name=Size,json=size" json:"Size"`
			// // -可选参数. 目前可用参数包括:
			// // -DynamicLoadNumber, 合法的数字型. 在库级别定义为REPO_LEVEL_ON_GPU时有效. 表明在启动时,按照数据新旧,最多加载到显存中的数据量
			// // -GPUThreads, [1,1,0,1]字符串格式. 在库级别定义为REPO_LEVEL_ON_GPU时必需. 表明数据在多个GPU上的分布. [1,1,0,1]表示当前服务器有三个GPU卡, 但是数据被平均存储在0,1和3号GPU卡上.

			Params: map[string]string{
				"GPUThreads":      repoParams.GPU,
				"UseFeatureIDMap": repoParams.UseFeatureIDMap,
				"PreFilter":       repoParams.PreFilter,
			},
		},
	}

	start := time.Now()

	for {
		log.Debugf("Sending request to rpc://%v//AddRepo, req = %v", addr, req)
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		defer cancel()

		resp, err := conn.client.RepoOperation(ctx, req)
		if err != nil || resp.Context.Status != FseReponseStatusOK {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/AddRepo\n%v\n%v", addr, err, resp)
			return errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/AddRepo, resp = %v, took: %v", addr, resp, time.Since(start))
		break
	}

	return nil
}

func (fps *FsePoolService) rankFeature(query *RankFeatureQuery) ([]*dg_model.RankItem, error) {
	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	var conn *fseConnection
	var addr string

	conn = fps.connectionPool[fps.getConnectIndex()]
	addr = fps.addr

	rankerContext, err := fps.newContext()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	params := make(map[string]string)

	if query.Confidence > 0 {
		params["ScoreThreshold"] = strconv.FormatFloat(float64(query.Confidence), 'f', -1, 32)
	}

	if query.TopX > 0 {
		max := strconv.Itoa(query.TopX)
		params["MaxCandidates"] = max
		params["PageSize"] = max // ranker默认100
	}

	if query.SensorIds != nil {
		params["Locations"] = strings.Join(query.SensorIds, ",")
	} else {
		params["Locations"] = "0"
	}

	if query.RecognizeType == dg_model.RecognizeType_REC_TYPE_FACE {
		params["Normalization"] = "true"
	}

	if query.RecognizeType == dg_model.RecognizeType_REC_TYPE_VEHICLE {
		params["Normalization"] = "false"
	}

	params["RepoId"] = query.RepoId
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

	req := &dg_model.RankFeatureRequest{
		Context: rankerContext,
		ObjectFeature: &dg_model.ObjectProperty{
			Feature: query.Feature,
		},

		Params: params,
	}

	start := time.Now()
	var resp *dg_model.RankFeatureResponse

	for {
		tryTimes++

		ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/RankFeature, req = %v", addr, req)

		resp, err = conn.client.RankFeature(ctx, req)
		if err != nil {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/RankFeature\n%v", addr, err)
			return nil, errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/RankFeature, resp = %v, took: %v", addr, resp, time.Since(start))
		break
	}

	return resp.Candidates, nil
}

func (fps *FsePoolService) rankFeatureWithContext(ctx context.Context, query *RankFeatureQuery) ([]*dg_model.RankItem, error) {
	fps.serviceOperationLock.Lock()
	defer fps.serviceOperationLock.Unlock()

	var tryTimes int
	var conn *fseConnection
	var addr string

	conn = fps.connectionPool[fps.getConnectIndex()]
	addr = fps.addr

	rankerContext, err := fps.newContext()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	params := make(map[string]string)

	if query.Confidence > 0 {
		params["ScoreThreshold"] = strconv.FormatFloat(float64(query.Confidence), 'f', -1, 32)
	}

	if query.TopX > 0 {
		max := strconv.Itoa(query.TopX)
		params["MaxCandidates"] = max
		params["PageSize"] = max // ranker默认100
	}

	if query.SensorIds != nil {
		params["Locations"] = strings.Join(query.SensorIds, ",")
	} else {
		params["Locations"] = "0"
	}

	if query.RecognizeType == dg_model.RecognizeType_REC_TYPE_FACE {
		params["Normalization"] = "true"
	}

	if query.RecognizeType == dg_model.RecognizeType_REC_TYPE_VEHICLE {
		params["Normalization"] = "false"
	}

	params["RepoId"] = query.RepoId
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

	req := &dg_model.RankFeatureRequest{
		Context: rankerContext,
		ObjectFeature: &dg_model.ObjectProperty{
			Feature: query.Feature,
		},

		Params: params,
	}

	start := time.Now()
	var resp *dg_model.RankFeatureResponse

	for {
		tryTimes++

		// ctx, cancel := context.WithTimeout(context.Background(), fps.timeoutSecond*time.Second)
		// defer cancel()

		log.Debugf("Sending request to rpc://%v/RankFeature, req = %v", addr, req)

		resp, err = conn.client.RankFeature(ctx, req)
		if err != nil {
			if tryTimes <= fps.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/RankFeature\n%v", addr, err)
			return nil, errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/RankFeature, resp = %v, took: %v", addr, resp, time.Since(start))
		break
	}

	return resp.Candidates, nil
}

func (fps *FsePoolService) getConnectIndex() int {
	fps.connectLock.Lock()
	defer fps.connectLock.Unlock()

	index := fps.connectIndex
	fps.connectIndex = (fps.connectIndex + 1) % fps.connectNubmer

	return index
}

func (fps *FsePoolService) newContext() (*dg_model.RankRequestContext, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &dg_model.RankRequestContext{
		SessionId: id.String(),
	}, nil
}
