package clients

import (
	"codes/gocodes/dg/models"

	dg_model "codes/gocodes/dg/dg.model"

	linq "github.com/ahmetb/go-linq"
	"github.com/pkg/errors"
	"go.uber.org/dig"
)

type VseMatrixClientComposedParams struct {
	dig.In

	VseFaceClient    VseMatrixClient `name:"VseFace"`
	VseVehicleClient VseMatrixClient `name:"VseVehicle"`
	VseAllObjClient  VseMatrixClient `name:"VseAllObj"`
}

type VseMatrixClientComposed struct {
	faceClient    VseMatrixClient
	vehicleClient VseMatrixClient
	allObjClient  VseMatrixClient
}

func NewVseMatrixClientComposed(params VseMatrixClientComposedParams) (*VseMatrixClientComposed, error) {
	if params.VseAllObjClient == nil || params.VseFaceClient == nil || params.VseVehicleClient == nil {
		return nil, errors.New("nil arguments")
	}

	return &VseMatrixClientComposed{
		faceClient:    params.VseFaceClient,
		vehicleClient: params.VseVehicleClient,
		allObjClient:  params.VseAllObjClient,
	}, nil
}

func (vmc *VseMatrixClientComposed) DetectMultiFacesForRankerFeature(img *models.ImageQuery) (*models.AllObject, error) {

	function := []int32{int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT)}

	// 大图逻辑
	bigRet, err := vmc.faceClient.RecognizeAllObject(img, dg_model.RecognizeType_REC_TYPE_FACE, DetectModeBigPic, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if bigRet != nil && len(bigRet.Faces) > 0 {
		return bigRet, nil
	}

	smallRet, err := vmc.faceClient.RecognizeAllObject(img, dg_model.RecognizeType_REC_TYPE_FACE, DetectModeSmallPic, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if smallRet != nil && len(smallRet.Faces) > 0 {
		return smallRet, nil
	}

	finalyRet, err := vmc.faceClient.RecognizeAllObject(img, dg_model.RecognizeType_REC_TYPE_FACE, DetectModeBigFirstSmallAfter, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if finalyRet != nil && len(finalyRet.Faces) > 0 {
		return finalyRet, nil
	}

	return nil, nil
}

func (vmc *VseMatrixClientComposed) DetectMultiVehicleForRankerFeature(img *models.ImageQuery) (*models.AllObject, error) {
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

	// 大图逻辑
	bigRet, err := vmc.vehicleClient.RecognizeAllObject(img, dg_model.RecognizeType_REC_TYPE_VEHICLE, DetectModeBigPic, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if bigRet != nil && (len(bigRet.Vehicles) > 0 || len(bigRet.NonmotorVehicles) > 0 || len(bigRet.Pedestrian) > 0) {
		return bigRet, nil
	}

	// 小图逻辑
	smallRet, err := vmc.vehicleClient.RecognizeAllObject(img, dg_model.RecognizeType_REC_TYPE_VEHICLE, DetectModeSmallPic, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if smallRet != nil && (len(smallRet.Vehicles) > 0 || len(smallRet.NonmotorVehicles) > 0 || len(smallRet.Pedestrian) > 0) {
		return smallRet, nil
	}

	// 先大图后小图
	finalyRet, err := vmc.vehicleClient.RecognizeAllObject(img, dg_model.RecognizeType_REC_TYPE_VEHICLE, DetectModeBigFirstSmallAfter, function)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if finalyRet != nil && (len(finalyRet.Vehicles) > 0 || len(finalyRet.NonmotorVehicles) > 0 || len(finalyRet.Pedestrian) > 0) {
		return smallRet, nil
	}

	return nil, nil
}

func (vmc *VseMatrixClientComposed) DetectObject(img *models.ImageQuery, recognizeType dg_model.RecognizeType, detectMode int) (*models.AllObject, error) {
	var ret *models.AllObject
	var err error

	function := make([]int32, 0)

	if recognizeType == dg_model.RecognizeType_REC_TYPE_FACE {
		function = append(function, int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT))

		ret, err = vmc.faceClient.RecognizeAllObject(img, recognizeType, detectMode, function)
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

		ret, err = vmc.vehicleClient.RecognizeAllObject(img, recognizeType, detectMode, function)
		if err != nil {
			return nil, errors.WithStack(err)
		}
	}

	return ret, nil
}

func (vmc *VseMatrixClientComposed) DetectMultiFaces(img *models.ImageQuery, detectMode int) ([]*models.ImageResult, error) {
	if img.Feature != "" {
		return nil, nil
	}

	recognizeType := dg_model.RecognizeType_REC_TYPE_FACE
	function := []int32{int32(dg_model.RecognizeFunctions_RECFUNC_FACE_DETECT)}

	ret, err := vmc.faceClient.RecognizeAllObject(img, recognizeType, detectMode, function)
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

func (vmc *VseMatrixClientComposed) GetBestFaceFeature(img *models.ImageQuery) (string, error) {
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

func (vmc *VseMatrixClientComposed) GetIndexFromFace(indexType dg_model.IndexType) (*dg_model.IndexTxtResponse, error) {
	return vmc.faceClient.GetIndexTxt(indexType)
}

func (vmc *VseMatrixClientComposed) GetIndexFromVehicle(indexType dg_model.IndexType) (*dg_model.IndexTxtResponse, error) {
	return vmc.vehicleClient.GetIndexTxt(indexType)
}
