package models

import (
	"fmt"
)

type FaceWhiteEventQuery struct {
	*CapturedFaceQuery
}

func (query *FaceWhiteEventQuery) Validate() error {
	if len(query.Images) > 0 {
		return fmt.Errorf("Query by image is not supported")
	}
	return nil
}
