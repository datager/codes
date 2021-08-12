package clients

import (
	"codes/gocodes/dg/models"

	linq "github.com/ahmetb/go-linq"
	"github.com/pkg/errors"
	"go.uber.org/dig"
	dg_model "codes/gocodes/dg/dg.model"
)

type VseMatrixHTTPComposedParams struct {
	dig.In

	VseFaceClient    *VseMatrixHTTPClient `name:"VseFace"`
	VseVehicleClient *VseMatrixHTTPClient `name:"VseVehicle"`
}

type VseMatrixHTTPComposed struct {
	faceClient    *VseMatrixHTTPClient
	vehicleClient *VseMatrixHTTPClient
}

func NewVseMatrixHTTPComposed(params VseMatrixHTTPComposedParams) (*VseMatrixHTTPComposed, error) {
	if params.VseFaceClient == nil || params.VseVehicleClient == nil {
		return nil, errors.New("nil arguments")
	}

	return &VseMatrixHTTPComposed{
		faceClient:    params.VseFaceClient,
		vehicleClient: params.VseVehicleClient,
	}, nil
}

const (
	FacePath    = "/vse/face/rec/image"
	VehiclePath = "/vse/vehicle/rec/image"
)

func (vmc *VseMatrixHTTPComposed) DetectMultiFacesForRankerFeature(img *models.ImageQuery) (*models.AllObject, error) {

	function := []int32{int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT)}
	facePath := FacePath

	type result struct {
		key string
		ret *models.AllObject
	}

	retChan := make(chan *result, 3)
	errChan := make(chan error, 3)
	defer func() {
		close(retChan)
		close(errChan)
	}()

	// 大图逻辑
	go func() {
		bigRet, err := vmc.faceClient.RecognizeAllObject(facePath, img, dg_model.RecognizeType_REC_TYPE_FACE, DetectModeBigPic, function)
		errChan <- err
		retChan <- &result{
			key: "big",
			ret: bigRet,
		}

	}()

	// 小图逻辑
	go func() {
		smallRet, err := vmc.faceClient.RecognizeAllObject(facePath, img, dg_model.RecognizeType_REC_TYPE_FACE, DetectModeSmallPic, function)
		errChan <- err
		retChan <- &result{
			key: "small",
			ret: smallRet,
		}
	}()

	// 先大图后小图
	go func() {
		finalyRet, err := vmc.faceClient.RecognizeAllObject(facePath, img, dg_model.RecognizeType_REC_TYPE_FACE, DetectModeBigFirstSmallAfter, function)
		errChan <- err
		retChan <- &result{
			key: "finally",
			ret: finalyRet,
		}
	}()

	var retErr error
	var res *models.AllObject
	for i := 0; i < 3; i++ {
		err := <-errChan
		if err != nil {
			retErr = errors.WithStack(err)
		}
	}

	var bigRet *models.AllObject
	var smallRet *models.AllObject
	var finallyRet *models.AllObject

	for i := 0; i < 3; i++ {
		ret := <-retChan
		if ret == nil {
			continue
		}

		if ret.ret == nil {
			continue
		}

		if ret.key == "big" {
			bigRet = ret.ret
		}

		if ret.key == "small" {
			smallRet = ret.ret
		}

		if ret.key == "finally" {
			finallyRet = ret.ret
		}
	}

	if bigRet == nil {
		return nil, errors.New("big picture model get nil result")
	}

	if smallRet == nil {
		return nil, errors.New("small picture model get nil result")
	}

	if finallyRet == nil {
		return nil, errors.New("model 2 get nil result")
	}

	if len(bigRet.Faces) == 1 && len(smallRet.Faces) == 0 {
		res = bigRet
	} else if len(bigRet.Faces) == 0 && len(smallRet.Faces) == 1 {
		res = smallRet
	} else if len(bigRet.Faces) > 1 && len(smallRet.Faces) > 1 {
		bigConfidence := bigRet.Faces[0].Confidence
		smallConfidence := smallRet.Faces[0].Confidence

		if bigConfidence > smallConfidence {
			res = bigRet
		} else {
			res = smallRet
		}
	} else {
		res = finallyRet
	}

	return res, errors.WithStack(retErr)
}

func (vmc *VseMatrixHTTPComposed) DetectMultiVehicleForRankerFeature(img *models.ImageQuery) (*models.AllObject, error) {
	function := make([]int32, 0)
	function = append(function,
		int32(dg_model.RecognizeFunctions_RECFUNC_VEHICLE_DETECT),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_STYLE),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_COLOR),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_MARKER),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_PLATE),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_FEATURE),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_DRIVER_BELT),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_DRIVER_PHONE),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_CODRIVER_BELT),
		// int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_DANGEROUS_ARTICLES_DETECT),
		// int32(dg_model.RecognizeFunctions_RECFUNC_PEDESTRIAN_ATTRIBUTE),
		// int32(dg_model.RecognizeFunctions_RECFUNC_PEDESTRIAN_FEATURE),
		// int32(dg_model.RecognizeFunctions_RECFUNC_NON_MOTOR_ATTRIBUTE),
	)

	vehiclePath := VehiclePath

	// 先大图后小图
	finalyRet, err := vmc.vehicleClient.RecognizeAllObject(vehiclePath, img, dg_model.RecognizeType_REC_TYPE_VEHICLE, DetectModeBigFirstSmallAfter, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return finalyRet, nil
}

func (vmc *VseMatrixHTTPComposed) DetectObject(img *models.ImageQuery, recognizeType dg_model.RecognizeType, detectMode int) (*models.AllObject, error) {
	var ret *models.AllObject
	var err error

	facePath := "/vse/face/rec/image"
	vehiclePath := "/vse/vehicle/rec/image"

	function := make([]int32, 0)

	if recognizeType == dg_model.RecognizeType_REC_TYPE_FACE {
		function = append(function, int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT))

		ret, err = vmc.faceClient.RecognizeAllObject(facePath, img, recognizeType, detectMode, function)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	if recognizeType == dg_model.RecognizeType_REC_TYPE_VEHICLE {
		function = append(function,
			int32(dg_model.RecognizeFunctions_RECFUNC_VEHICLE_DETECT),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_STYLE),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_COLOR),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_MARKER),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_PLATE),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_FEATURE),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_DRIVER_BELT),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_DRIVER_PHONE),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_CODRIVER_BELT),
			int32(dg_model.RecognizeFunctions_RECFUNC_MOTOR_DANGEROUS_ARTICLES_DETECT),
			int32(dg_model.RecognizeFunctions_RECFUNC_PEDESTRIAN_ATTRIBUTE),
			int32(dg_model.RecognizeFunctions_RECFUNC_PEDESTRIAN_FEATURE),
			int32(dg_model.RecognizeFunctions_RECFUNC_NON_MOTOR_ATTRIBUTE),
		)

		ret, err = vmc.vehicleClient.RecognizeAllObject(vehiclePath, img, recognizeType, detectMode, function)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return ret, nil
}

func (vmc *VseMatrixHTTPComposed) DetectMultiFaces(img *models.ImageQuery, detectMode int) ([]*models.ImageResult, error) {
	if img.Feature != "" {
		return nil, nil
	}

	recognizeType := dg_model.RecognizeType_REC_TYPE_FACE
	function := []int32{int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT)}

	facePath := "/vse/face/rec/image"

	ret, err := vmc.faceClient.RecognizeAllObject(facePath, img, recognizeType, detectMode, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if len(ret.Faces) == 0 {
		return nil, errors.New("Faces not found")
	}

	results := make([]*models.ImageResult, 0)
	for _, v := range ret.Faces {
		if v.QualityOK {
			results = append(results, v)
		}
	}

	return results, nil
}

func (vmc *VseMatrixHTTPComposed) GetBestFaceFeature(img *models.ImageQuery) (string, error) {
	if img.Feature != "" {
		return img.Feature, nil
	}
	faces, err := vmc.DetectMultiFaces(img, DetectModeBigFirstSmallAfter)
	if err != nil {
		return "", errors.WithStack(err)
	}
	if len(faces) == 0 {
		return "", errors.New("Faces not found")
	}
	face := linq.From(faces).Where(func(f interface{}) bool {
		return f.(*models.ImageResult).QualityOK
	}).OrderByDescending(func(f interface{}) interface{} {
		return f.(*models.ImageResult).Confidence
	}).First()
	if face == nil {
		return "", errors.New("Failed to get best face")
	}
	return (face.(*models.ImageResult)).Feature, nil
}

func (vmc *VseMatrixHTTPComposed) GetIndexFromFace(indexType dg_model.IndexType) (*dg_model.IndexTxtResponse, error) {
	path := "/vse/vehicle/rec/index/txt"
	return vmc.faceClient.GetIndexTxt(path, indexType)
}

func (vmc *VseMatrixHTTPComposed) GetIndexFromVehicle(indexType dg_model.IndexType) (*dg_model.IndexTxtResponse, error) {
	path := "/vse/vehicle/rec/index/txt"
	return vmc.vehicleClient.GetIndexTxt(path, indexType)
}
