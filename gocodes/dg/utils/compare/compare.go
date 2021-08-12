package compare

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"math"

	"codes/gocodes/dg/utils/log"
)

type Similarity struct {
	IsOpen                  bool
	CosThreashold           float32
	CosAParams              []float32
	CosBParams              []float32
	ThresParams             []float32
	LinearAlpha             float32
	LinearBeta              float32
	SlidingWindowMillSecond int64
	ConfigType              int
	TopNum                  int
}

func PairScore(firstFeature, secondFeature string, similarity *Similarity) (float32, error) {
	if firstFeature == "" || secondFeature == "" || similarity == nil {
		return 0, errors.New("feature is nil")
	}

	score := CosVerification(
		DecodeFeature(firstFeature), DecodeFeature(secondFeature),
		similarity.CosAParams,
		similarity.CosBParams,
		similarity.ThresParams,
		similarity.LinearAlpha,
		similarity.LinearBeta,
	)

	return score, nil
}

func DecodeFeature(featureStr string) []float32 {
	decodeBytes, err := base64.StdEncoding.DecodeString(featureStr)
	if err != nil {
		log.Errorf("Feature decode error: %s", featureStr)
		return nil
	}

	featureArr := make([]float32, len(decodeBytes)/4)

	for i := 0; i < len(decodeBytes); i += 4 {
		bits := binary.LittleEndian.Uint32(decodeBytes[i : i+4])
		float := math.Float32frombits(bits)
		featureArr[i/4] = float
	}

	return featureArr
}

func CosVerification(firtVector []float32, secondeVector []float32, cosABelow []float32, cosBBelow []float32, cosThreds []float32, linearAlpha float32, linearBeta float32) float32 {

	if len(firtVector) != len(secondeVector) {
		return 0
	}

	var dot float32
	for i := 0; i < len(firtVector); i++ {
		dot += firtVector[i] * secondeVector[i]
	}

	originalScore := 0.5 + 0.5*dot
	if originalScore > 1 {
		originalScore = 1
	} else if originalScore < 0 {
		originalScore = 0
	}

	expScore := originalScore
	var nonlinearScore float32

	for i := 0; i < len(cosABelow); i++ {
		if i == len(cosABelow)-1 {
			nonlinearScore = cosABelow[i]*expScore + cosBBelow[i]
			break
		}
		if expScore < cosThreds[i] {
			nonlinearScore = cosABelow[i]*expScore + cosBBelow[i]
			break
		}
	}

	score := linearAlpha*nonlinearScore + linearBeta

	if score > 1 {
		return 1
	}

	return score
}
