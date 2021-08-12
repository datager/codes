package json

import (
	"encoding/json"
)

var (
	//json          = jsoniter.ConfigCompatibleWithStandardLibrary
	Marshal       = json.Marshal
	MarshalIndent = json.MarshalIndent
	NewDecoder    = json.NewDecoder
	Unmarshal     = json.Unmarshal
)
