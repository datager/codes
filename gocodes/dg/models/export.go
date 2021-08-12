package models

import (
	"image"
	"net/http"
	"time"

	"codes/gocodes/dg/utils"
)

const (
	ExportTaskStatus_Created = iota
	ExportTaskStatus_Pending
	ExportTaskStatus_Running
	ExportTaskStatus_RanToCompletion
	ExportTaskStatus_Faulted
	ExportTaskStatus_Canceled
)

const (
	ExportTaskType_Unknown = iota
	ExportTaskType_Captured
	ExportTaskType_Civil
	ExportTaskType_Event
	ExportTaskType_FailedCivil
	ExportTaskType_WhiteEvent
	ExportTaskType_PersonFile
)

var (
	ExportTaskTypes = []int{
		ExportTaskType_Unknown,
		ExportTaskType_Captured,
		ExportTaskType_Civil,
		ExportTaskType_Event,
		ExportTaskType_FailedCivil,
		ExportTaskType_WhiteEvent,
		ExportTaskType_PersonFile,
	}
)

type ExportTaskStatus int

type ExportRequest struct {
	Type           int
	Query          interface{}
	FileNamePrefix string
}

type ExportResponse struct {
	Total     int
	Processed int
	Status    ExportTaskStatus
	StartAt   int64
	UpdateAt  int64
}

func (resp *ExportResponse) Start() {
	resp.Status = ExportTaskStatus_Running
	now := utils.GetNowTs()
	resp.StartAt = now
	resp.UpdateAt = now
}

func (resp *ExportResponse) SetProcessed(processed int) {
	resp.Processed = processed
	resp.UpdateAt = utils.GetNowTs()
}

type ExportTask struct {
	Id                       string
	Request                  *ExportRequest
	Response                 *ExportResponse
	Timestamp                int64
	TmpPath                  string        `json:"-"` // json ignore
	DownloadAssetTimeout     time.Duration `json:"-"` // json ignore
	DownloadAssetConcurrency int           `json:"-"` // json ignore
	HTTPClient               *http.Client
}

type ExportAsset struct {
	Url    string
	SaveAs string
	Config *image.Config
}

type ExportRow struct {
	Value  []string
	Assets map[int]*ExportAsset
}
