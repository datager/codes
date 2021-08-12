package models

const (
	KindOrg = iota + 1
	KindSensor
	KindFeature
)

const (
	ActionAdd = iota + 1
	ActionUpdate
	ActionDelete
	ActionSyncAll
)

type FerryMessage struct {
	Topic     string
	Partition int
	Key       string
	Data      string
}

type BaseData struct {
	Kind     int
	MetaData []byte
	Action   int
}

type SyncOrgsRet struct {
	ErrorCode int
	Message   string
}
