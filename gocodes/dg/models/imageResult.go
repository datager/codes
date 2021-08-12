package models

type ImageResult struct {
	ImageUri          string
	ThumbnailImageUri string
	CutboardImageUri  string
	CutboardX         int
	CutboardY         int
	CutboardWidth     int
	CutboardHeight    int
	CutboardResWidth  int
	CutboardResHeight int
	CutboardBinData   string
	Confidence        float32
	BinData           string
	Feature           string
	IsRanked          int
	QualityOK         bool
	Status            TaskStatus
	ImageId           string
	ImageType         int
	PlateText         string
}

type AllObject struct {
	Faces            []*ImageResult
	Vehicles         []*ImageResult
	Pedestrian       []*ImageResult
	NonmotorVehicles []*ImageResult
}

type DetectObjectRet map[int][]*ImageResult
