package clients

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"

// 	"codes/gocodes/dg/models"
// 	"codes/gocodes/dg/utils"
// 	"codes/gocodes/dg/utils/config"
// 	"codes/gocodes/dg/utils/log"

// 	dg_model "codes/gocodes/dg/dg.model"

// 	linq "github.com/ahmetb/language-linq"
// 	"github.com/gofrs/uuid"
// 	"github.com/golang/glog"
// 	"github.com/pkg/errors"
// 	"google.golang.org/grpc"
// )

// var matrix_service *MatrixPoolService
// var once sync.Once

// type connection struct {
// 	client dg_model.WitnessServiceClient
// 	conn   *grpc.ClientConn
// }

// type matrixConfig struct {
// 	RfcnAddr             string
// 	AlignAddr            string
// 	FeatureAddr          string
// 	RfcnContext          dg_model.WitnessRequestContext
// 	AlignContext         dg_model.WitnessRequestContext
// 	FeatureContext       dg_model.WitnessRequestContext
// 	RfcnConnectNumber    int
// 	AlignConnectNumber   int
// 	FeatureConnectNumber int
// 	TimeoutSecond        uint
// 	TryTimes             uint
// }

// func getMatrixConfig(conf config.Config) *matrixConfig {
// 	rfcnAddr := conf.GetString("services.matrix_rfcn.addr")
// 	alignAddr := conf.GetString("services.matrix_fcn.addr")
// 	featureAddr := conf.GetString("services.matrix_feature.addr")

// 	if rfcnAddr == "" || alignAddr == "" || featureAddr == "" {
// 		panic("addr is empty")
// 	}

// 	rfcnConnectNumber := conf.GetIntOrDefault("services.maxtri_pool.rfcnConnectNumber", 5)
// 	alignConnectNumber := conf.GetIntOrDefault("services.maxtri_pool.alignConnectNumber", 5)
// 	featureConnectNumber := conf.GetIntOrDefault("services.maxtri_pool.featureConnectNumber", 5)

// 	timeoutSecond := conf.GetIntOrDefault("services.maxtri_pool.timeoutSecond", 5)
// 	tryTimes := conf.GetIntOrDefault("services.maxtri_pool.tryTimes", 3)

// 	return &matrixConfig{
// 		RfcnAddr:             rfcnAddr,
// 		AlignAddr:            alignAddr,
// 		FeatureAddr:          featureAddr,
// 		RfcnConnectNumber:    rfcnConnectNumber,
// 		AlignConnectNumber:   alignConnectNumber,
// 		FeatureConnectNumber: featureConnectNumber,
// 		TimeoutSecond:        uint(timeoutSecond),
// 		TryTimes:             uint(tryTimes),
// 		RfcnContext: dg_model.WitnessRequestContext{
// 			SessionId: "rfcn",
// 			Functions: []int32{200},
// 			Type:      2,
// 		},
// 		AlignContext: dg_model.WitnessRequestContext{
// 			SessionId: "fcn_align",
// 			Functions: []int32{201, 202, 203},
// 			Type:      2,
// 		},
// 		FeatureContext: dg_model.WitnessRequestContext{
// 			SessionId: "feature",
// 			Functions: []int32{204},
// 			Type:      2,
// 		},
// 	}
// }

// type MatrixPoolService struct {
// 	config                  *matrixConfig
// 	rfcnConnectionPool      []*connection
// 	alignConnectionPool     []*connection
// 	featureConnectionPool   []*connection
// 	serviceOperationLock    sync.RWMutex
// 	rfcnConnectIndex        int
// 	rfcnConnectIndexLock    sync.Mutex
// 	alignConnectIndex       int
// 	alignConnectIndexLock   sync.Mutex
// 	featureConnectIndex     int
// 	isServiceOpen           bool
// 	featureConnectIndexLock sync.Mutex
// }

// func NewMatrixPoolService(conf config.Config) *MatrixPoolService {

// 	once.Do(func() {
// 		matrix_service = &MatrixPoolService{
// 			config: getMatrixConfig(conf),
// 		}

// 	})
// 	return matrix_service

// }

// func (this *MatrixPoolService) Start() {
// 	this.serviceOperationLock.Lock()
// 	defer this.serviceOperationLock.Unlock()

// 	if this.isServiceOpen {
// 		return
// 	}

// 	// init rfcn connect pool
// 	for {

// 		if len(this.rfcnConnectionPool) < this.config.RfcnConnectNumber {
// 			rfcnConn, err := grpc.Dial(this.config.RfcnAddr, grpc.WithInsecure())
// 			if err != nil {
// 				glog.Errorln("[CROATIA_SERVICE_MATRIX] rfcn connect err:", err)
// 				<-time.After(time.Second)
// 				continue
// 			}
// 			rfcnClient := dg_model.NewWitnessServiceClient(rfcnConn)
// 			this.rfcnConnectionPool = append(this.rfcnConnectionPool, &connection{
// 				conn:   rfcnConn,
// 				client: rfcnClient,
// 			})
// 			glog.Infoln("[CROATIA_SERVICE_MATRIX] create a new connect to rfcn,total connect num is", len(this.rfcnConnectionPool))
// 			continue
// 		}

// 		break
// 	}

// 	//init align connect pool
// 	for {

// 		if len(this.alignConnectionPool) < this.config.AlignConnectNumber {
// 			alignConn, err := grpc.Dial(this.config.AlignAddr, grpc.WithInsecure())
// 			if err != nil {
// 				glog.Errorln("[CROATIA_SERVICE_MATRIX] algin connect err:", err)
// 				<-time.After(time.Second)
// 				continue
// 			}
// 			alignClient := dg_model.NewWitnessServiceClient(alignConn)
// 			this.alignConnectionPool = append(this.alignConnectionPool, &connection{
// 				conn:   alignConn,
// 				client: alignClient,
// 			})
// 			glog.Infoln("[CROATIA_SERVICE_MATRIX] create a new connect to align,total connect num is", len(this.alignConnectionPool))
// 			continue
// 		}

// 		break
// 	}

// 	//init feature connect pool
// 	for {

// 		if len(this.featureConnectionPool) < this.config.FeatureConnectNumber {
// 			featureConn, err := grpc.Dial(this.config.FeatureAddr, grpc.WithInsecure())
// 			if err != nil {
// 				glog.Errorln("[CROATIA_SERVICE_MATRIX] algin connect err:", err)
// 				<-time.After(time.Second)
// 				continue
// 			}
// 			featureClient := dg_model.NewWitnessServiceClient(featureConn)
// 			this.featureConnectionPool = append(this.featureConnectionPool, &connection{
// 				conn:   featureConn,
// 				client: featureClient,
// 			})
// 			glog.Infoln("[CROATIA_SERVICE_MATRIX] create a new connect to feature,total connect num is", len(this.featureConnectionPool))
// 			continue
// 		}

// 		glog.Infoln("[CROATIA_SERVICE_MATRIX] matrix service start")

// 		break
// 	}

// 	this.isServiceOpen = true

// }

// func (this *MatrixPoolService) Stop() {
// 	this.serviceOperationLock.Lock()
// 	defer this.serviceOperationLock.Unlock()

// 	if !this.isServiceOpen {
// 		return
// 	}

// 	//close rfcn connect pool
// 	for _, connection := range this.rfcnConnectionPool {
// 		connection.client = nil
// 		connection.conn.Close()
// 	}

// 	this.rfcnConnectionPool = []*connection{}

// 	//close algin connect pool
// 	for _, connection := range this.alignConnectionPool {
// 		connection.client = nil
// 		connection.conn.Close()
// 	}

// 	this.alignConnectionPool = []*connection{}

// 	//close feature connect pool
// 	for _, connection := range this.featureConnectionPool {
// 		connection.client = nil
// 		connection.conn.Close()
// 	}

// 	this.featureConnectionPool = []*connection{}

// 	this.isServiceOpen = false

// 	glog.Infoln("[CROATIA_SERVICE_MATRIX] matrix service stop")
// }

// //get rfcn response
// func (this *MatrixPoolService) Rfcn(request *dg_model.WitnessRequest) (*dg_model.WitnessResponse, error) {

// 	this.serviceOperationLock.RLock()
// 	defer this.serviceOperationLock.RUnlock()

// 	if this.config.RfcnConnectNumber == 0 || !this.isServiceOpen {
// 		return nil, errors.New("[CROATIA_SERVICE_MATRIX] rfcn connect pool is empty")
// 	}

// 	var tryTimes uint = 0
// 	var err error
// 	var rfcnResponse *dg_model.WitnessResponse

// 	//get rfcn client index
// 	rfcnClientiIndex := this.getRfcnConnectIndex()

// 	for {
// 		tryTimes++
// 		//send post request
// 		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(this.config.TimeoutSecond)*time.Second)
// 		defer cancel()
// 		rfcnResponse, err = this.rfcnConnectionPool[rfcnClientiIndex].client.Recognize(ctx, request)
// 		if err != nil {
// 			if tryTimes <= this.config.TryTimes {
// 				continue
// 			}
// 			return nil, errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX],rfcn err:", err))
// 		}
// 		break
// 	}

// 	return rfcnResponse, nil
// }

// //get align response
// func (this *MatrixPoolService) FcnAlign(request *dg_model.WitnessRequest) (*dg_model.WitnessResponse, error) {
// 	this.serviceOperationLock.RLock()
// 	defer this.serviceOperationLock.RUnlock()

// 	if this.config.AlignConnectNumber == 0 || !this.isServiceOpen {
// 		return nil, errors.New("[CROATIA_SERVICE_MATRIX] align connect pool is empty")
// 	}

// 	var tryTimes uint = 0
// 	var err error
// 	var alignResponse *dg_model.WitnessResponse

// 	//get algin client index
// 	alignClientiIndex := this.getAlginConnectIndex()

// 	for {
// 		tryTimes++
// 		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(this.config.TimeoutSecond)*time.Second)
// 		defer cancel()
// 		alignResponse, err = this.alignConnectionPool[alignClientiIndex].client.Recognize(ctx, request)
// 		if err != nil {
// 			if tryTimes <= this.config.TryTimes {
// 				continue
// 			}
// 			return nil, errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX],align err:", err))

// 		}
// 		break
// 	}

// 	return alignResponse, nil

// }

// //get feature response
// func (this *MatrixPoolService) Feature(request *dg_model.WitnessRequest) (*dg_model.WitnessResponse, error) {

// 	this.serviceOperationLock.RLock()
// 	defer this.serviceOperationLock.RUnlock()

// 	if this.config.FeatureConnectNumber == 0 || !this.isServiceOpen {
// 		return nil, errors.New("[CROATIA_SERVICE_MATRIX] feature connect pool is empty")
// 	}

// 	var tryTimes uint = 0
// 	var err error
// 	var featureResponse *dg_model.WitnessResponse

// 	//get algin client index
// 	featureClientiIndex := this.getFeatureConnectIndex()

// 	for {
// 		tryTimes++
// 		//send post request
// 		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(this.config.TimeoutSecond)*time.Second)
// 		defer cancel()
// 		featureResponse, err = this.featureConnectionPool[featureClientiIndex].client.Recognize(ctx, request)
// 		if err != nil {
// 			if tryTimes <= this.config.TryTimes {
// 				continue
// 			}
// 			return nil, errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX],feature err:", err))

// 		}
// 		break

// 	}
// 	return featureResponse, nil
// }

// func (this *MatrixPoolService) getRfcnConnectIndex() int {
// 	this.rfcnConnectIndexLock.Lock()
// 	defer this.rfcnConnectIndexLock.Unlock()
// 	index := this.rfcnConnectIndex
// 	this.rfcnConnectIndex = (this.rfcnConnectIndex + 1) % this.config.RfcnConnectNumber
// 	return index
// }

// func (this *MatrixPoolService) getAlginConnectIndex() int {
// 	this.alignConnectIndexLock.Lock()
// 	defer this.alignConnectIndexLock.Unlock()
// 	index := this.alignConnectIndex
// 	this.alignConnectIndex = (this.alignConnectIndex + 1) % this.config.AlignConnectNumber
// 	return index
// }

// func (this *MatrixPoolService) getFeatureConnectIndex() int {
// 	this.featureConnectIndexLock.Lock()
// 	defer this.featureConnectIndexLock.Unlock()
// 	index := this.featureConnectIndex
// 	this.featureConnectIndex = (this.featureConnectIndex + 1) % this.config.FeatureConnectNumber
// 	return index
// }

// func (this *MatrixPoolService) Getfeature(imageUrl string, bindata string) (string, error) {

// 	this.serviceOperationLock.RLock()
// 	defer this.serviceOperationLock.RUnlock()

// 	rfcnContext := this.config.RfcnContext
// 	alignContext := this.config.AlignContext
// 	featureContext := this.config.FeatureContext

// 	witnessImage := dg_model.WitnessImage{
// 		Data: &dg_model.Image{
// 			BinData: bindata,
// 			URI:     imageUrl,
// 		},
// 		UserObject: []*dg_model.WitnessUserObject{
// 			&dg_model.WitnessUserObject{},
// 		},
// 	}

// 	req := &dg_model.WitnessRequest{
// 		Context: &rfcnContext,
// 		Image:   &witnessImage,
// 	}

// 	//rfcn
// 	resp, err := this.Rfcn(req)
// 	if err != nil {
// 		return "", errors.WithStack(err)
// 	}

// 	if resp == nil ||
// 		resp.GetResult() == nil ||
// 		resp.GetResult().GetFaces() == nil ||
// 		len(resp.GetResult().GetFaces()) == 0 {
// 		// if no faces --> fcn & align
// 		req.Context = &alignContext
// 		for _, function := range this.config.RfcnContext.Functions {
// 			req.Context.Functions = append(req.Context.Functions, function)
// 		}
// 	} else {

// 		// detect multi face return error
// 		if len(resp.GetResult().GetFaces()) > 1 {
// 			return "", fmt.Errorf("multi face return error")
// 		}

// 		//check face rect
// 		err = this.checkFaceRect(resp, resp.GetResult().GetFaces()[0])
// 		if err != nil {
// 			return "", errors.WithStack(err)
// 		}

// 		req.Context = &alignContext
// 		faces := resp.GetResult().GetFaces()
// 		req.GetImage().GetUserObject()[0].Type = dg_model.WitnessUserObjectType_WITNESS_USER_OBJECT_FACE
// 		req.GetImage().GetUserObject()[0].RotatedRect = faces[0].GetImg().GetRect()
// 	}

// 	//algin or fcn&align
// 	resp, err = this.FcnAlign(req)
// 	if err != nil {
// 		return "", errors.WithStack(err)
// 	}

// 	if resp == nil ||
// 		resp.GetResult() == nil ||
// 		resp.GetResult().GetFaces() == nil ||
// 		len(resp.GetResult().GetFaces()) == 0 {
// 		return "", fmt.Errorf("no face error")

// 	} else {

// 		// detect multi face return error
// 		if len(resp.GetResult().GetFaces()) > 1 {
// 			return "", fmt.Errorf("multi face error")
// 		}

// 		//check face rect
// 		err = this.checkFaceRect(resp, resp.GetResult().GetFaces()[0])
// 		if err != nil {
// 			return "", errors.WithStack(err)
// 		}

// 		req.Context = &featureContext
// 		faces := resp.GetResult().GetFaces()
// 		req.GetImage().GetUserObject()[0].AlignResult = faces[0].GetAlignResult()
// 		req.GetImage().GetUserObject()[0].Type = dg_model.WitnessUserObjectType_WITNESS_USER_OBJECT_FACE
// 	}

// 	//quality
// 	err = this.checkQuality(resp.GetResult().GetFaces()[0])
// 	if err != nil {
// 		return "", errors.WithStack(err)
// 	}

// 	//feature
// 	resp, err = this.Feature(req)
// 	if err != nil {
// 		return "", errors.WithStack(err)
// 	}

// 	var feature string
// 	if resp == nil ||
// 		resp.GetResult() == nil ||
// 		resp.GetResult().GetFaces() == nil ||
// 		len(resp.GetResult().GetFaces()) == 0 ||
// 		resp.GetResult().GetFaces()[0].GetFeatures() == "" {

// 		return "", fmt.Errorf("feature is empty")

// 	} else {
// 		feature = resp.GetResult().GetFaces()[0].GetFeatures()
// 	}

// 	return feature, nil
// }

// func (this *MatrixPoolService) checkQuality(face *dg_model.RecFace) error {
// 	alignScore := face.Qualities["AlignScore"]
// 	pitch := face.Qualities["Pitch"]
// 	yaw := face.Qualities["Yaw"]

// 	if alignScore < 0.05 {
// 		return errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX] quality err: alignscore:", alignScore, "threadhold:0.005"))
// 	}

// 	if alignScore > 0.01 {
// 		if pitch < -45 || pitch > 45 || yaw < -45 || yaw > 45 {
// 			return errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX] quality err: pitch:", pitch, " yaw:", yaw))
// 		}
// 	}

// 	return nil

// }

// func (this *MatrixPoolService) checkFaceRect(response *dg_model.WitnessResponse, face *dg_model.RecFace) error {

// 	//get src image width and height
// 	if response == nil || response.Result == nil || response.Result.Image == nil ||
// 		response.Result.Image.Data == nil || response.Result.Image.Data.Width == 0 ||
// 		response.Result.Image.Data.Height == 0 {
// 		return errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX] matrix src image data is nil"))
// 	}

// 	imageWidth := response.Result.Image.Data.Width
// 	imageHeight := response.Result.Image.Data.Height

// 	//get face image width and height
// 	if face == nil || face.Img == nil || face.Img.Rect == nil {
// 		return errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX] matrix detect result image data is nil"))
// 	}

// 	rectWidth := face.Img.Rect.Width
// 	rectHeight := face.Img.Rect.Height

// 	//get rate
// 	widthRate := float32(rectWidth) / float32(imageWidth)
// 	heightRate := float32(rectHeight) / float32(imageHeight)

// 	if widthRate <= 0.005 || heightRate <= 0.005 {
// 		return errors.New(fmt.Sprintln("[CROATIA_SERVICE_MATRIX] matrix detect result image cover src image rate below threashold,width rate:", widthRate, " height rate:", heightRate))
// 	}

// 	return nil

// }

// func (this *MatrixPoolService) GetBestFaceFeature(img *models.ImageQuery) (string, error) {
// 	if img.Feature != "" {
// 		return img.Feature, nil
// 	}
// 	faces, err := this.RecognizeMultiFace(img)
// 	if err != nil {
// 		return "", errors.WithStack(err)
// 	}
// 	if len(faces) == 0 {
// 		return "", errors.New("Faces not found")
// 	}
// 	face := linq.From(faces).Where(func(f interface{}) bool {
// 		return f.(*models.ImageResult).QualityOK
// 	}).OrderByDescending(func(f interface{}) interface{} {
// 		return f.(*models.ImageResult).Confidence
// 	}).First()
// 	if face == nil {
// 		return "", errors.New("Failed to get best face")
// 	}
// 	return (face.(*models.ImageResult)).Feature, nil
// }

// func (this *MatrixPoolService) RecognizeMultiFace(img *models.ImageQuery) ([]*models.ImageResult, error) {
// 	// rfcn
// 	rfcnReq, err := this.buildRFCNRequest(img)
// 	if nil != err {
// 		return nil, errors.Wrap(err, "build RFCN request failed")
// 	}
// 	rfcnResp, err := this.Rfcn(rfcnReq)
// 	if nil != err {
// 		return nil, errors.Wrap(err, "RFCN failed")
// 	}
// 	fcnReq, err := this.buildFCNRequest(img, rfcnResp)
// 	if nil != err {
// 		return nil, errors.Wrap(err, "build FCN request failed")
// 	}
// 	fcnResp, err := this.FcnAlign(fcnReq)
// 	if nil != err {
// 		return nil, errors.Wrap(err, "FCN failed")
// 	}
// 	featureReq, err := this.buildFeatureRequest(img, rfcnResp, fcnResp)
// 	if nil != err {
// 		return nil, errors.Wrap(err, "build feature request failed")
// 	}
// 	featureResp, err := this.Feature(featureReq)
// 	if nil != err {
// 		return nil, errors.Wrap(err, "feature failed")
// 	}
// 	results := featureResp.GetResult()
// 	if results == nil {
// 		return nil, errors.New("Got empty results from Feature")
// 	}
// 	faces := results.GetFaces()
// 	rets := make([]*models.ImageResult, len(faces))
// 	for i, face := range results.GetFaces() {
// 		rets[i] = &models.ImageResult{
// 			CutboardX:         int(face.GetImg().Cutboard.GetX()),
// 			CutboardY:         int(face.GetImg().Cutboard.GetY()),
// 			CutboardWidth:     int(face.GetImg().GetCutboard().GetWidth()),
// 			CutboardHeight:    int(face.GetImg().GetCutboard().GetHeight()),
// 			CutboardResWidth:  int(face.GetImg().GetCutboard().GetResWidth()),
// 			CutboardResHeight: int(face.GetImg().GetCutboard().GetResHeight()),
// 			Confidence:        face.GetConfidence(),
// 			Feature:           face.GetFeatures(),
// 			QualityOK:         true,
// 		}
// 	}
// 	return rets, nil
// }

// func (this *MatrixPoolService) buildBaseRequest(img *models.ImageQuery) (*dg_model.WitnessRequest, error) {
// 	id, err := uuid.NewV4()
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}
// 	return &dg_model.WitnessRequest{
// 		Context: &dg_model.WitnessRequestContext{
// 			Type:      int32(dg_model.RecognizeType_REC_TYPE_FACE),
// 			SessionId: id.String(),
// 		},
// 		Image: &dg_model.WitnessImage{
// 			Data: &dg_model.Image{
// 				URI:     img.Url,
// 				BinData: img.BinData,
// 			},
// 			WitnessMetaData: &dg_model.SrcMetadata{
// 				Timestamp: utils.GetNowTs(),
// 				ObjType:   dg_model.ObjType_OBJ_TYPE_FACE,
// 			},
// 		},
// 	}, nil
// }

// func (this *MatrixPoolService) buildRFCNRequest(img *models.ImageQuery) (*dg_model.WitnessRequest, error) {
// 	req, err := buildBaseRequest(img)
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}
// 	req.Context.Functions = []int32{int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT)}
// 	return req, nil
// }

// func (this *MatrixPoolService) buildFCNRequest(img *models.ImageQuery, respRFCN *dg_model.WitnessResponse) (*dg_model.WitnessRequest, error) {
// 	results := respRFCN.GetResult()
// 	if results == nil {
// 		return nil, errors.New("Got empty results from RFCN")
// 	}
// 	req, err := buildBaseRequest(img)
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}
// 	if results == nil {
// 		req.Context.Functions = []int32{
// 			int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT),
// 			int32(dg_model.RecognizeFunctions_RECFUNC_FACE_ALIGNMENT),
// 			int32(dg_model.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV1),
// 			int32(dg_model.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV2)}
// 	} else {
// 		req.Context.Functions = []int32{
// 			int32(dg_model.RecognizeFunctions_RECFUNC_FACE_ALIGNMENT),
// 			int32(dg_model.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV1),
// 			int32(dg_model.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV2)}
// 		faces := results.GetFaces()
// 		for _, face := range faces {
// 			userObj := &dg_model.WitnessUserObject{
// 				Type:        dg_model.WitnessUserObjectType_WITNESS_USER_OBJECT_FACE,
// 				RotatedRect: face.Img.Rect,
// 			}
// 			req.Image.UserObject = append(req.Image.UserObject, userObj)
// 		}
// 	}
// 	return req, nil
// }

// func (this *MatrixPoolService) buildFeatureRequest(img *models.ImageQuery, respRFCN *dg_model.WitnessResponse, respFCN *dg_model.WitnessResponse) (*dg_model.WitnessRequest, error) {
// 	results := respFCN.GetResult()
// 	if results == nil {
// 		return nil, errors.New("Got empty results from FCN")
// 	}
// 	req, err := buildFCNRequest(img, respRFCN)
// 	if err != nil {
// 		return nil, errors.WithStack(err)
// 	}
// 	if faces := results.GetFaces(); len(faces) > 0 {
// 		for index := 0; index < len(faces); index++ {
// 			if faces[index].AlignResult == nil {
// 				log.Warnln(fmt.Sprintf("Got empty align result at index %d, ignore", index))
// 				continue
// 			} else if len(req.Image.UserObject) <= index {
// 				userObj := &dg_model.WitnessUserObject{
// 					Type:        dg_model.WitnessUserObjectType_WITNESS_USER_OBJECT_FACE,
// 					RotatedRect: faces[index].Img.Rect,
// 					AlignResult: faces[index].AlignResult,
// 				}
// 				req.Image.UserObject = append(req.Image.UserObject, userObj)
// 			} else {
// 				req.Image.UserObject[index].AlignResult = faces[index].AlignResult
// 			}
// 		}
// 	}
// 	req.Context.Functions = []int32{
// 		int32(dg_model.RecognizeFunctions_RECFUNC_FACE_FEATURE),
// 		int32(dg_model.RecognizeFunctions_RECFUNC_FACE_ATTRIBUTE),
// 	}
// 	return req, nil
// }
