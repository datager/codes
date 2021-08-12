package models

type VehicleCaptureIndex struct {
	Timestamp    int64
	RepoID       string
	SensorID     string
	SensorName   string
	VehicleID    string
	VehicleReID  string
	VehicleVid   string
	ImageResult  ImageResult
	PlateText    string
	Speed        int32
	Direction    int32
	FaceID       string
	BrandID      int32
	SubBrandID   int32
	ModelYearID  int32
	TypeID       int32
	Side         int32
	ColorID      int32
	PlateTypeID  int32
	PlateColorID int32
	Symbols      SymbolSlice
	Specials     []int64
	Lane         string
	HasFace      bool
	SymbolsDesc  string
}

type VehiclesIndexExtend struct {
	VehicleCapture
	Sensor Sensor
}

func (ve VehiclesIndexExtend) GetSensor() *Sensor {
	return &ve.Sensor
}

func (ve VehiclesIndexExtend) GetCapturedTimestamp() int64 {
	return ve.VehicleCapture.Timestamp
}

type VehicleAnalyzeData struct {
	VehicleVid  string
	VehicleReID string
	SensorID    string
	Timestamp   int64
	Longitude   float64
	Latitude    float64
}
