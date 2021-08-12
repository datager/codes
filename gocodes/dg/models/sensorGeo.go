package models

import (
	"fmt"
	"math"
)

const (
	DefaultSensorLongitude = -200
	DefaultSensorLatitude  = -200
)

type SensorGeo struct {
	Longitude float64
	Latitude  float64
}

func (geo *SensorGeo) Validate() error {
	if math.Abs(float64(geo.Longitude)) > 180 {
		return fmt.Errorf("Invalid longitude %f", geo.Longitude)
	}
	if math.Abs(float64(geo.Latitude)) > 90 {
		return fmt.Errorf("Invalid latitude %f", geo.Latitude)
	}
	return nil
}
