package clients

import (
	"codes/gocodes/dg/models"
)

type ElasticsearchClient interface {
	GetFaceCount(request models.FaceConditionRequest) (int64, error)
	GetPedestrianCount(request models.PedestrianConditionRequest) (int64, error)
	GetVehicleCount(request models.VehicleConditionRequest) (int64, error)
	GetNonmotorCount(request models.NonmotorConditionRequest) (int64, error)
}
