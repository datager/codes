package wolf

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/service/login"
	"codes/gocodes/dg/utils/clients"
	"codes/gocodes/dg/utils/json"

	"github.com/golang/glog"
	jsoniter "github.com/json-iterator/go"
)

// 提前准备测试数据
/*
delete from faces_index where ts >= 1585031485359 and ts < 1585117885359;
INSERT INTO ""("ts", "sensor_id", "sensor_int_id", "org_code", "face_id", "face_reid", "face_vid", "relation_types", "gender_id", "age_id", "nation_id", "glass_id", "mask_id", "hat_id", "halmet_id", "cutboard_image_uri") VALUES (1585110044679, '656aa5de-17e8-4aa7-9770-02af56d022bc', 44, 468750000, '9d55931b-c159-44bf-a246-106d92543423', 'aa7683f1-4494-48a8-9f71-62a19f030f1e', '8062c929-77d8-4402-9248-cd82ac690bd2', 0, 2, 28, 1, 2, 2, 1, 1, 'http://192.168.2.142:8501/api/v2/file/39/ae416f49881daa');
INSERT INTO ""("ts", "sensor_id", "sensor_int_id", "org_code", "face_id", "face_reid", "face_vid", "relation_types", "gender_id", "age_id", "nation_id", "glass_id", "mask_id", "hat_id", "halmet_id", "cutboard_image_uri") VALUES (1585107833923, '656aa5de-17e8-4aa7-9770-02af56d022bc', 44, 468750000, '9c4c1091-c29b-44da-8492-cebb3129fe45', '656d6ddb-c450-4ce5-88d0-9f295f149821', '8062c929-77d8-4402-9248-cd82ac690bd2', 0, 2, 31, 1, 2, 2, 1, 1, 'http://192.168.2.142:8501/api/v2/file/36/ae1bfa255829a7');
INSERT INTO ""("ts", "sensor_id", "sensor_int_id", "org_code", "face_id", "face_reid", "face_vid", "relation_types", "gender_id", "age_id", "nation_id", "glass_id", "mask_id", "hat_id", "halmet_id", "cutboard_image_uri") VALUES (1585107803844, '656aa5de-17e8-4aa7-9770-02af56d022bc', 44, 468750000, '7066d254-325d-41a1-b92c-b359928e3781', '9d3e1314-3739-4e63-b0cc-1093936a4979', '8062c929-77d8-4402-9248-cd82ac690bd2', 0, 2, 28, 1, 2, 2, 1, 1, 'http://192.168.2.142:8501/api/v2/file/41/ae1b5e1582747f');
INSERT INTO ""("ts", "sensor_id", "sensor_int_id", "org_code", "face_id", "face_reid", "face_vid", "relation_types", "gender_id", "age_id", "nation_id", "glass_id", "mask_id", "hat_id", "halmet_id", "cutboard_image_uri") VALUES (1585107537625, '656aa5de-17e8-4aa7-9770-02af56d022bc', 44, 468750000, '107b3cf8-b8c2-4a55-9e9b-65063bf63077', '4ec3120b-bf46-4425-8eb8-8a804d4ec60e', '8062c929-77d8-4402-9248-cd82ac690bd2', 0, 2, 31, 1, 2, 2, 1, 1, 'http://192.168.2.142:8501/api/v2/file/39/ae16f1908ef628');
*/

// 因为只有4个ts, 所以
// ts1=1585107537625
// ts2=1585107803844
// ts3=1585107833923
// ts4=1585110044679

var ( // limit=0即为全部, 在sql里不拼offset和limit的条件
	faceCaptureControllerTest     *FaceCaptureControllerSortAndFromPageTest
	faceCaptureControllerTestOnce sync.Once
)

type FaceCaptureControllerSortAndFromPageTest struct {
	c *clients.HTTPClient

	testReq models.FaceConditionRequest

	testFuncs []func(bool, bool, int, int) (*models.SearchResult, error)
}

func InitFaceCaptureControllerTest() *FaceCaptureControllerSortAndFromPageTest {
	token, err := login.Login()
	if err != nil {
		glog.Fatal(err)
	}

	faceCaptureControllerTestOnce.Do(func() {
		faceCaptureControllerTest = new(FaceCaptureControllerSortAndFromPageTest)
	})

	faceCaptureControllerTest.c = clients.NewHTTPClient("http://192.168.2.142:3000/api", time.Duration(1000000)*time.Second, nil)
	faceCaptureControllerTest.c.SetHeader(http.Header{"Content-Type": []string{"application/json"}})
	faceCaptureControllerTest.c.SetHeader(http.Header{"authorization": []string{"Bearer " + token}})

	faceCaptureControllerTest.testReq = models.FaceConditionRequest{
		Vid:     "8062c929-77d8-4402-9248-cd82ac690bd2",
		OrderBy: "ts",
		Time: models.StartAndEndTimestamp{
			StartTimestamp: 1585031485359,
			EndTimestamp:   1585117885359,
		},
	}

	faceCaptureControllerTest.testFuncs = append(faceCaptureControllerTest.testFuncs,
		faceCaptureControllerTest.testFaceOrderAsc,
	)

	return faceCaptureControllerTest
}

func (cct *FaceCaptureControllerSortAndFromPageTest) TestRun() {
	for funcNo := range cct.testFuncs {
		var err error
		glog.Info("----------1-----------\n降序 首页, 期望直接按降序查, 期望返回pg1为ts4, ts3, pg2为ts2, ts1")
		_, err = cct.testFuncs[funcNo](false, false, 0, 2)
		if err != nil {
			glog.Error(err)
		}

		_, err = cct.testFuncs[funcNo](false, false, 2, 2)
		if err != nil {
			glog.Error(err)
		}

		glog.Info("----------2-----------\n降序 末页, 期望直接按(!降序)查, 再reverse, 期望返回pg1为ts2, ts1, pg2为ts4, ts3")
		_, err = cct.testFuncs[funcNo](false, true, 0, 2)
		if err != nil {
			glog.Error(err)
		}

		_, err = cct.testFuncs[funcNo](false, true, 2, 2)
		if err != nil {
			glog.Error(err)
		}

		glog.Info("----------3-----------\n升序 首页, 期望直接按升序查, 期望返回pg1为ts1, ts2, pg2为ts3, ts4")
		_, err = cct.testFuncs[funcNo](true, false, 0, 2)
		if err != nil {
			glog.Error(err)
		}

		_, err = cct.testFuncs[funcNo](true, false, 2, 2)
		if err != nil {
			glog.Error(err)
		}

		glog.Info("----------4-----------\n升序 末页, 期望直接按(!升序)查, 再reverse, 期望返回pg1为ts3, ts4, pg2为ts1, ts2")
		_, err = cct.testFuncs[funcNo](true, true, 0, 2)
		if err != nil {
			glog.Error(err)
		}

		_, err = cct.testFuncs[funcNo](true, true, 2, 2)
		if err != nil {
			glog.Error(err)
		}
	}
	fmt.Print("done")
}

func (cct *FaceCaptureControllerSortAndFromPageTest) testFaceOrderAsc(asc, tailPage bool, offset, limit int) (*models.SearchResult, error) {
	reqInner := cct.testReq
	reqInner.OrderAsc, reqInner.FromTailPage = asc, tailPage
	reqInner.Offset, reqInner.Limit = offset, limit

	reqInnerBytes, err := jsoniter.Marshal(reqInner)
	if err != nil {
		return nil, err
	}
	glog.Info("req ", reqInner.OrderAsc, reqInner.FromTailPage, reqInner.Offset, reqInner.Limit)

	reqOuter := &models.SearchListRequest{
		Type:    models.TypeSearchCondition,
		Payload: reqInnerBytes,
	}
	reqOuterBytes, err := jsoniter.Marshal(reqOuter)
	if err != nil {
		return nil, err
	}

	respByte, err := cct.c.Post("/face/captures", bytes.NewReader(reqOuterBytes))
	if err != nil {
		return nil, err
	}

	resp := new(models.SearchResult)
	err = json.Unmarshal(respByte, &resp)
	if err != nil {
		return nil, err
	}

	glog.Info("resp ", resp.NextOffset)
	for _, r := range resp.Rets {
		mdl := new(models.CapturedFace)
		b, _ := jsoniter.Marshal(r)
		_ = jsoniter.Unmarshal(b, mdl)

		switch mdl.Timestamp {
		case 1585107537625:
			glog.Info("ts1")
		case 1585107803844:
			glog.Info("ts2")
		case 1585107833923:
			glog.Info("ts3")
		case 1585110044679:
			glog.Info("ts4")
		default:
			glog.Info("bad case")
		}
	}

	return resp, nil
}
