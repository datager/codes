package wolf

import (
	"bytes"
	"codes/gocodes/dg/configs"
	"codes/gocodes/dg/dbengine"
	"codes/gocodes/dg/entities"
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/service/login"
	"codes/gocodes/dg/utils/uuid"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func InsertMockFaceRepos() {
	stAll := time.Now()
	threadNum := configs.GetInstance().FaceRepoInsertConfig.ThreadNum
	glog.Infof("开始插入库测试数据 start InsertMockFaceRepos ThreadNum:%v", threadNum)

	token, err := login.Login()
	if err != nil {
		panic(err)
	}

	g, _ := errgroup.WithContext(context.TODO())
	for i := 0; i < int(threadNum); i++ {
		req := &models.SetFaceRepoRequest{
			//RepoID:       ,
			OrgID:        "库并发测试",
			RepoName:     fmt.Sprintf("库并发测试%v-%v", uuid.NewV4().String(), i),
			Comment:      "库并发测试",
			NameListAttr: 1,
			SystemTag:    1,
			Type:         1,
		}

		g.Go(func() error {
			return insertFaceRepo(req, token)
		})
	}

	if err := g.Wait(); err != nil {
		glog.Fatal(err)
	}

	glog.Infof("结束插入库测试数据(并行) finish InsertMockFaceRepos cost %v", time.Since(stAll))
}

func DeleteMockedFaceRepos() {
	token, err := login.Login()
	if err != nil {
		panic(err)
	}

	st := time.Now()
	glog.Infoln("开始清理库内测试数据 start DeleteMockFaceRepos")

	mockedVehicleRepos := queryVehicleFaceRepos()
	for _, repo := range mockedVehicleRepos {
		glog.Infof("delete repo %v", repo)
		err := deleteRepo(repo.RepoId, token)
		if err != nil {
			panic(err)
		}
	}

	glog.Infof("结束清理库内测试数据(串行) finish DeleteMockFaceRepos len %v, cost %v", len(mockedVehicleRepos), time.Since(st))
}

func insertFaceRepo(param *models.SetFaceRepoRequest, token string) error {
	st := time.Now()
	reqJs, err := json.Marshal(param)
	if err != nil {
		return err
	}

	r := bytes.NewReader(reqJs)
	req, err := http.NewRequest(http.MethodPost, models.FaceRepoCreateURL, r)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode/100 != 2 { // should be 204
		return err
	}

	glog.Infof("insertFaceRepo cost %v, %v", time.Since(st), string(reqJs))
	return nil
}

//func queryVehicleFaceRepos(token string) []*models.FaceRepo {
//	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%v?RepoName=库并发测试&Limit=1000&Offset=0&Type=1", models.FaceRepoGetURL), nil)
//	if err != nil {
//		panic(err)
//	}
//
//	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
//	resp, err := http.DefaultClient.Do(req)
//	if err != nil {
//		panic(err)
//	}
//	defer resp.Body.Close()
//
//	bodyBytes, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		panic(err)
//	}
//
//	ret := &models.FaceRepoGetResult{}
//	err = json.Unmarshal(bodyBytes, &ret)
//	if err != nil {
//		panic(err)
//	}
//
//	return ret.Rets
//}

func queryVehicleFaceRepos() []*entities.FaceRepos {
	ret := make([]*entities.FaceRepos, 0)
	err := dbengine.GetV5Instance().Where("repo_name like '%库并发测试%'").Where("type = ?", 1).Find(&ret)
	if err != nil {
		panic(err)
	}
	return ret
}

func deleteRepo(repoID, token string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%v/%v", models.FaceRepoDeleteURL, repoID), nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("authorization", fmt.Sprintf("Bearer %v", token))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode/100 != 2 { // should be 204
		return err
	}
	return nil
}
