package clients

import (
	"fmt"

	"go.uber.org/dig"

	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/log"

	linq "github.com/ahmetb/go-linq"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	globalmodels "codes/gocodes/dg/dg.model"
)

type MatrixClientComposedParams struct {
	dig.In

	ClientRFCN    MatrixClient `name:"MatrixRFCN"`
	ClientFCN     MatrixClient `name:"MatrixFCN"`
	ClientFeature MatrixClient `name:"MatrixFeature"`
}

type MatrixClientComposed struct {
	clientRFCN    MatrixClient
	clientFCN     MatrixClient
	clientFeature MatrixClient
}

func NewMatrixClientComposed(params MatrixClientComposedParams) (*MatrixClientComposed, error) {
	if params.ClientRFCN == nil || params.ClientFCN == nil || params.ClientFeature == nil {
		return nil, errors.New("Nil arguments")
	}
	return &MatrixClientComposed{
		clientRFCN:    params.ClientRFCN,
		clientFCN:     params.ClientFCN,
		clientFeature: params.ClientFeature,
	}, nil
}

func (this *MatrixClientComposed) GetBestFaceFeature(img *models.ImageQuery) (string, error) {
	if img.Feature != "" {
		return img.Feature, nil
	}
	faces, err := this.RecognizeMultiFace(img)
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

func (this *MatrixClientComposed) RecognizeMultiFace(img *models.ImageQuery) ([]*models.ImageResult, error) {
	// rfcn
	rfcnReq, err := buildRFCNRequest(img)
	if nil != err {
		return nil, errors.Wrap(err, "build RFCN request failed")
	}
	rfcnResp, err := this.clientRFCN.BatchRecognize(rfcnReq)
	if nil != err {
		return nil, errors.Wrap(err, "RFCN failed")
	}
	fcnReq, err := buildFCNRequest(img, rfcnResp)
	if nil != err {
		return nil, errors.Wrap(err, "build FCN request failed")
	}
	fcnResp, err := this.clientFCN.BatchRecognize(fcnReq)
	if nil != err {
		return nil, errors.Wrap(err, "FCN failed")
	}
	featureReq, err := buildFeatureRequest(img, rfcnResp, fcnResp)
	if nil != err {
		return nil, errors.Wrap(err, "build feature request failed")
	}
	featureResp, err := this.clientFeature.BatchRecognize(featureReq)
	if nil != err {
		return nil, errors.Wrap(err, "feature failed")
	}
	results := featureResp.GetResult()
	if results == nil {
		return nil, errors.New("Got empty results from Feature")
	}
	faces := results.GetFaces()
	rets := make([]*models.ImageResult, len(faces))
	for i, face := range results.GetFaces() {
		rets[i] = &models.ImageResult{
			CutboardX:         int(face.GetImg().Cutboard.GetX()),
			CutboardY:         int(face.GetImg().Cutboard.GetY()),
			CutboardWidth:     int(face.GetImg().GetCutboard().GetWidth()),
			CutboardHeight:    int(face.GetImg().GetCutboard().GetHeight()),
			CutboardResWidth:  int(face.GetImg().GetCutboard().GetResWidth()),
			CutboardResHeight: int(face.GetImg().GetCutboard().GetResHeight()),
			Confidence:        face.GetConfidence(),
			Feature:           face.GetFeatures(),
			QualityOK:         true,
		}
	}
	return rets, nil
}

func buildBaseRequest(img *models.ImageQuery) (*globalmodels.WitnessRequest, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &globalmodels.WitnessRequest{
		Context: &globalmodels.WitnessRequestContext{
			Type:      int32(globalmodels.RecognizeType_REC_TYPE_FACE),
			SessionId: id.String(),
		},
		Image: &globalmodels.WitnessImage{
			Data: &globalmodels.Image{
				URI:     img.Url,
				BinData: img.BinData,
			},
			WitnessMetaData: &globalmodels.SrcMetadata{
				Timestamp: utils.GetNowTs(),
				ObjType:   globalmodels.ObjType_OBJ_TYPE_FACE,
			},
		},
	}, nil
}

func buildRFCNRequest(img *models.ImageQuery) (*globalmodels.WitnessRequest, error) {
	req, err := buildBaseRequest(img)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Context.Functions = []int32{int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_DETECT)}
	return req, nil
}

func buildFCNRequest(img *models.ImageQuery, respRFCN *globalmodels.WitnessResponse) (*globalmodels.WitnessRequest, error) {
	results := respRFCN.GetResult()
	if results == nil {
		return nil, errors.New("Got empty results from RFCN")
	}
	req, err := buildBaseRequest(img)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if results == nil {
		req.Context.Functions = []int32{
			int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_DETECT),
			int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_ALIGNMENT),
			int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV1),
			int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV2)}
	} else {
		req.Context.Functions = []int32{
			int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_ALIGNMENT),
			int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV1),
			int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_QUALITY_LV2)}
		faces := results.GetFaces()
		for _, face := range faces {
			userObj := &globalmodels.WitnessUserObject{
				Type:        globalmodels.WitnessUserObjectType_WITNESS_USER_OBJECT_FACE,
				RotatedRect: face.Img.Rect,
			}
			req.Image.UserObject = append(req.Image.UserObject, userObj)
		}
	}
	return req, nil
}

func buildFeatureRequest(img *models.ImageQuery, respRFCN *globalmodels.WitnessResponse, respFCN *globalmodels.WitnessResponse) (*globalmodels.WitnessRequest, error) {
	results := respFCN.GetResult()
	if results == nil {
		return nil, errors.New("Got empty results from FCN")
	}
	req, err := buildFCNRequest(img, respRFCN)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if faces := results.GetFaces(); len(faces) > 0 {
		for index := 0; index < len(faces); index++ {
			if faces[index].AlignResult == nil {
				log.Warnln(fmt.Sprintf("Got empty align result at index %d, ignore", index))
				continue
			} else if len(req.Image.UserObject) <= index {
				userObj := &globalmodels.WitnessUserObject{
					Type:        globalmodels.WitnessUserObjectType_WITNESS_USER_OBJECT_FACE,
					RotatedRect: faces[index].Img.Rect,
					AlignResult: faces[index].AlignResult,
				}
				req.Image.UserObject = append(req.Image.UserObject, userObj)
			} else {
				req.Image.UserObject[index].AlignResult = faces[index].AlignResult
			}
		}
	}
	req.Context.Functions = []int32{
		int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_FEATURE),
		int32(globalmodels.RecognizeFunctions_RECFUNC_FACE_ATTRIBUTE),
	}
	return req, nil
}
