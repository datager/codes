package clients

import (
	"context"
	"fmt"
	"sync"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/log"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	dg_model "codes/gocodes/dg/dg.model"

	"google.golang.org/grpc"
)

const (
	DetectModeBigPic             = 0
	DetectModeSmallPic           = 1
	DetectModeBigFirstSmallAfter = 2
)

const (
	VseResponseStatusOK = "200"
)

type vseConnection struct {
	client dg_model.WitnessServiceClient
	conn   *grpc.ClientConn
}
type VseMatrixPoolService struct {
	addr          string
	connectNubmer int

	timeoutSecond time.Duration
	tryTimes      int

	connectionPool []*vseConnection

	serviceOperationLock sync.RWMutex
	connectIndex         int
	connectLock          sync.Mutex

	isServiceOpen bool
}

func NewVseMatrixPoolService(addr string, connectNumber int, timeoutSecond time.Duration, tryTimes int) *VseMatrixPoolService {
	if addr == "" {
		panic("addr is empty")
	}

	service := &VseMatrixPoolService{
		addr:          addr,
		connectNubmer: connectNumber,
		timeoutSecond: timeoutSecond,
		tryTimes:      tryTimes,
	}

	service.init()

	return service
}

func (vmp *VseMatrixPoolService) init() {
	vmp.serviceOperationLock.Lock()
	defer vmp.serviceOperationLock.Unlock()

	if vmp.isServiceOpen {
		return
	}

	for {

		if len(vmp.connectionPool) < vmp.connectNubmer {
			faceConn, err := grpc.Dial(vmp.addr, grpc.WithInsecure())
			if err != nil {
				log.Errorln("face connect err:", err)
				<-time.After(time.Second)
				continue
			}

			faceClient := dg_model.NewWitnessServiceClient(faceConn)
			vmp.connectionPool = append(vmp.connectionPool, &vseConnection{
				conn:   faceConn,
				client: faceClient,
			})

			// log.Infoln("create a new connect to vse face, total connect num is ", len(vmp.faceConnectionPool))
			continue
		}

		break
	}

	vmp.isServiceOpen = true
}

func (vmp *VseMatrixPoolService) Close() error {
	vmp.serviceOperationLock.Lock()
	defer vmp.serviceOperationLock.Unlock()

	if !vmp.isServiceOpen {
		return nil
	}

	for _, connection := range vmp.connectionPool {
		connection.client = nil
		connection.conn.Close()
	}

	vmp.connectionPool = []*vseConnection{}

	vmp.isServiceOpen = false
	log.Infoln("vse matrix service closed")

	return nil
}

// func (vmp *VseMatrixPoolService) RecognizeFaces(img *models.ImageQuery, detectMode int) ([]*models.ImageResult, error) {
// 	if img.Feature != "" {
// 		return nil, nil
// 	}

// 	recognizeType := dg_model.RecognizeType_REC_TYPE_FACE
// 	objType := dg_model.ObjType_OBJ_TYPE_FACE
// 	function := []int32{int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT)}

// 	ret, err := vmp.recognizeAllObject(img, recognizeType, detectMode, objType, function)
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}

// 	if len(ret.Faces) == 0 {
// 		return nil, errors.New("Faces not found")
// 	}

// 	results := make([]*models.ImageResult, 0)
// 	for _, v := range ret.Faces {
// 		if v.QualityOK {
// 			results = append(results, v)
// 		}
// 	}

// 	return results, nil
// }

func (vmp *VseMatrixPoolService) RecognizeAllObject(img *models.ImageQuery, recognizeType dg_model.RecognizeType, detectMode int, recognizeFunction []int32) (*models.AllObject, error) {
	vmp.serviceOperationLock.Lock()
	defer vmp.serviceOperationLock.Unlock()

	var tryTimes int
	var resp *dg_model.WitnessResponse

	req, err := vmp.buildRequest(img, recognizeType, detectMode, recognizeFunction)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	client := vmp.connectionPool[vmp.getConnectIndex()]

	start := time.Now()

	for {

		tryTimes++
		ctx, cancel := context.WithTimeout(context.Background(), vmp.timeoutSecond*time.Second)
		defer cancel()

		log.Debugf("Sending request to rpc://%v/RecognizeAllObject, req = %v", vmp.addr, req)

		resp, err = client.client.Recognize(ctx, req)
		if err != nil || resp.Context.Status != VseResponseStatusOK {
			if tryTimes <= vmp.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/RecognizeAllObject\n%+v", vmp.addr, err)

			return nil, fmt.Errorf("recognize face err:%v", err)
		}

		log.Debugf("Got response from rpc://%v/RecognizeAllObject, resp = %v, took: %v", vmp.addr, resp, time.Since(start))
		break
	}

	results := resp.GetResult()
	if results == nil {
		return nil, errors.New("Got empty results from feature")
	}

	faces := results.GetFaces()
	nonMotorVehicles := results.GetNonMotorVehicles()
	pedestrian := results.GetPedestrian()
	vehicles := results.GetVehicles()

	facesRet := make([]*models.ImageResult, len(faces))
	for i, face := range faces {
		facesRet[i] = &models.ImageResult{
			CutboardX:         int(face.GetImg().Cutboard.GetX()),
			CutboardY:         int(face.GetImg().Cutboard.GetY()),
			CutboardWidth:     int(face.GetImg().GetCutboard().GetWidth()),
			CutboardHeight:    int(face.GetImg().GetCutboard().GetHeight()),
			CutboardResWidth:  int(face.GetImg().GetCutboard().GetResWidth()),
			CutboardResHeight: int(face.GetImg().GetCutboard().GetResHeight()),
			Confidence:        face.GetImg().GetCutboard().GetConfidence(),
			Feature:           face.GetFeatures(),
			QualityOK:         true,
		}
	}

	nonMotorVehiclesRet := make([]*models.ImageResult, len(nonMotorVehicles))
	for i, v := range nonMotorVehicles {
		nonMotorVehiclesRet[i] = &models.ImageResult{
			CutboardX:         int(v.GetImg().Cutboard.GetX()),
			CutboardY:         int(v.GetImg().Cutboard.GetY()),
			CutboardWidth:     int(v.GetImg().GetCutboard().GetWidth()),
			CutboardHeight:    int(v.GetImg().GetCutboard().GetHeight()),
			CutboardResWidth:  int(v.GetImg().GetCutboard().GetResWidth()),
			CutboardResHeight: int(v.GetImg().GetCutboard().GetResHeight()),
			Confidence:        v.GetImg().GetCutboard().GetConfidence(),
			Feature:           v.GetFeatures(),
			QualityOK:         true,
		}
	}

	pedestrianRet := make([]*models.ImageResult, len(pedestrian))
	for i, v := range pedestrian {
		pedestrianRet[i] = &models.ImageResult{
			CutboardX:         int(v.GetImg().Cutboard.GetX()),
			CutboardY:         int(v.GetImg().Cutboard.GetY()),
			CutboardWidth:     int(v.GetImg().GetCutboard().GetWidth()),
			CutboardHeight:    int(v.GetImg().GetCutboard().GetHeight()),
			CutboardResWidth:  int(v.GetImg().GetCutboard().GetResWidth()),
			CutboardResHeight: int(v.GetImg().GetCutboard().GetResHeight()),
			Confidence:        v.GetImg().GetCutboard().GetConfidence(),
			Feature:           v.GetFeatures(),
			QualityOK:         true,
		}
	}

	vehiclesRet := make([]*models.ImageResult, len(vehicles))
	for i, v := range vehicles {
		vehiclesRet[i] = &models.ImageResult{
			CutboardX:         int(v.GetImg().Cutboard.GetX()),
			CutboardY:         int(v.GetImg().Cutboard.GetY()),
			CutboardWidth:     int(v.GetImg().GetCutboard().GetWidth()),
			CutboardHeight:    int(v.GetImg().GetCutboard().GetHeight()),
			CutboardResWidth:  int(v.GetImg().GetCutboard().GetResWidth()),
			CutboardResHeight: int(v.GetImg().GetCutboard().GetResHeight()),
			Confidence:        v.GetImg().GetCutboard().GetConfidence(),
			Feature:           v.GetFeatures(),
			QualityOK:         true,
		}
	}

	return &models.AllObject{
		Faces:            facesRet,
		Vehicles:         vehiclesRet,
		NonmotorVehicles: nonMotorVehiclesRet,
		Pedestrian:       pedestrianRet,
	}, nil
}

func (vmp *VseMatrixPoolService) GetIndexTxt(indexType dg_model.IndexType) (*dg_model.IndexTxtResponse, error) {
	vmp.serviceOperationLock.Lock()
	defer vmp.serviceOperationLock.Unlock()

	var tryTimes int
	var resp *dg_model.IndexTxtResponse
	var err error

	req := &dg_model.IndexTxtRequest{
		IndexType: indexType,
	}

	conn := vmp.connectionPool[vmp.getConnectIndex()]

	start := time.Now()

	for {
		tryTimes++

		// ctx, cancel := context.WithTimeout(context.Background(), vmp.timeoutSecond*time.Hour)
		// defer cancel()
		ctx := context.Background()

		log.Debugf("Sending request to rpc://%v/GetIndexTxt, req = %v", vmp.addr, req)

		log.Debugln("开始发送请求。。。")
		resp, err = conn.client.GetIndexTxt(ctx, req)
		log.Debugln("收到返回信息")
		if err != nil {
			if tryTimes <= vmp.tryTimes {
				continue
			}

			log.Errorf("Failed to send request to rpc://%v/GetIndexTxt\n%+v", vmp.addr, err)

			return nil, errors.WithStack(err)
		}

		log.Debugf("Got response from rpc://%v/GetIndexTxt, resp = %v, took: %v", vmp.addr, resp, time.Since(start))
		break
	}

	return resp, nil
}

func (vmp *VseMatrixPoolService) buildRequest(img *models.ImageQuery, recognizeType dg_model.RecognizeType, detectMode int, recognizeFunction []int32) (*dg_model.WitnessRequest, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	params := make(map[string]string)
	params["detect_mode"] = fmt.Sprintf("%v", detectMode)

	req := &dg_model.WitnessRequest{
		Context: &dg_model.WitnessRequestContext{
			Type:      int32(recognizeType),
			SessionId: id.String(),
			Params:    params,
			Functions: recognizeFunction,
		},
		Image: &dg_model.WitnessImage{
			Data: &dg_model.Image{
				URI:     img.Url,
				BinData: img.BinData,
			},
		},
	}

	return req, nil
}

func (vmp *VseMatrixPoolService) getConnectIndex() int {
	vmp.connectLock.Lock()
	defer vmp.connectLock.Unlock()

	index := vmp.connectIndex
	vmp.connectIndex = (vmp.connectIndex + 1) % vmp.connectNubmer

	return index
}
