package clients

import (
	"fmt"
	"time"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/json"

	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	dg_model "codes/gocodes/dg/dg.model"
)

type VseMatrixHTTPClient struct {
	*HTTPClient
}

func NewVseMatrixHTTPClient(addr string, timeoutSecond time.Duration) *VseMatrixHTTPClient {
	ret := &VseMatrixHTTPClient{
		HTTPClient: NewHTTPClient(addr, timeoutSecond, nil),
	}

	ret.HTTPClient.defaultRespReader = ret.readResponse

	return ret
}

func (vhc *VseMatrixHTTPClient) RecognizeAllObject(path string, img *models.ImageQuery, recognizeType dg_model.RecognizeType, detectMode int, recognizeFunction []int32) (*models.AllObject, error) {
	req, err := vhc.buildRequest(img, recognizeType, detectMode, recognizeFunction)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	response, err := vhc.PostJSON(path, req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var resp dg_model.WitnessResponse
	err = json.Unmarshal(response, &resp)
	if err != nil {
		return nil, errors.WithStack(err)
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
		var plateText string
		plates := v.GetPlates()
		if len(plates) > 0 {
			plateText = plates[0].GetPlateText()
		}
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
			PlateText:         plateText,
		}
	}

	return &models.AllObject{
		Faces:            facesRet,
		Vehicles:         vehiclesRet,
		NonmotorVehicles: nonMotorVehiclesRet,
		Pedestrian:       pedestrianRet,
	}, nil
}

func (vhc *VseMatrixHTTPClient) GetIndexTxt(path string, indexType dg_model.IndexType) (*dg_model.IndexTxtResponse, error) {
	req := dg_model.IndexTxtRequest{
		IndexType: indexType,
	}

	response, err := vhc.PostJSON(path, &req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var resp dg_model.IndexTxtResponse
	err = json.Unmarshal(response, &resp)

	return &resp, errors.WithStack(err)
}

func (vhc *VseMatrixHTTPClient) buildRequest(img *models.ImageQuery, recognizeType dg_model.RecognizeType, detectMode int, recognizeFunction []int32) (*dg_model.WitnessRequest, error) {
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
