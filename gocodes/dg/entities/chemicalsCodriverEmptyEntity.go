package entities

const ChemicalsCodriverEmptyTableName = "chemicals_codriver_empty_captures"

type ChemicalsCodriverEmptyEntity struct {
	Timestamp int64  `xorm:"ts"`
	CaptureId string `xorm:"capture_id"`
	SensorId  string `xorm:"sensor_id"`
	PlateText string `xorm:"plate_text"`
}

func (ChemicalsCodriverEmptyEntity) TableName() string {
	return ChemicalsCodriverEmptyTableName
}
