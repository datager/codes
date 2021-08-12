package models

type FaceEventAuditRequest struct {
	ItemMap map[string]int // 1:有效,2:无效,3:未知
	Comment string
}
