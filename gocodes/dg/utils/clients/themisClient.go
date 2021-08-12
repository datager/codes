package clients

import (
	"encoding/json"
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/config"
	"time"
)

type ThemisClient struct {
	*HTTPClient
}

func NewThemisClient(conf config.Config) *ThemisClient {
	timeoutSecond := conf.GetIntOrDefault("services.themis.timeout_second", 60)
	client := ThemisClient{
		HTTPClient: NewHTTPClient(conf.GetString("services.themis.addr"), time.Duration(timeoutSecond)*time.Second, nil),
	}
	return &client
}

func (themisClient ThemisClient) UnBeltedList(request *models.VehicleSpecialWithDriverAndPassengerRequest) (*models.ThemisResponse, error) {
	bodyBytes, err := themisClient.HTTPClient.PostJSON("/specials/unbeltedlist", request)
	if err != nil {
		return nil, err
	}
	var result = new(models.ThemisResponse)
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (themisClient ThemisClient) DriverCallList(request *models.VehicleSpecialRequest) (*models.ThemisResponse, error) {
	bodyBytes, err := themisClient.HTTPClient.PostJSON("/specials/drivecalllist", request)
	if err != nil {
		return nil, err
	}
	var result = new(models.ThemisResponse)
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (themisClient ThemisClient) CoverFaceList(request *models.VehicleSpecialWithDriverAndPassengerRequest) (*models.ThemisResponse, error) {
	bodyBytes, err := themisClient.HTTPClient.PostJSON("/specials/coverfacelist", request)
	if err != nil {
		return nil, err
	}
	var result = new(models.ThemisResponse)
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (themisClient ThemisClient) DriverSmokeList(request *models.VehicleSpecialRequest) (*models.ThemisResponse, error) {
	bodyBytes, err := themisClient.HTTPClient.PostJSON("/specials/drivesmokinglist", request)
	if err != nil {
		return nil, err
	}
	var result = new(models.ThemisResponse)
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (themisClient ThemisClient) ChemicalsCodriverEmptyList(request *models.VehicleSpecialRequest) (*models.ThemisResponse, error) {
	bodyBytes, err := themisClient.HTTPClient.PostJSON("/specials/chemicalscodriveremptylist", request)
	if err != nil {
		return nil, err
	}
	var result = new(models.ThemisResponse)
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return nil, err
	}
	return result, nil
}

func (themisClient ThemisClient) NonmotorWithPeopleList(request *models.NonmotorWithPeople) (*models.NonmotorThemisResponse, error) {
	bodyBytes, err := themisClient.HTTPClient.PostJSON("/specials/nonmotorwithpeoplelist", request)
	if err != nil {
		return nil, err
	}
	var result = new(models.NonmotorThemisResponse)
	if err := json.Unmarshal(bodyBytes, result); err != nil {
		return nil, err
	}
	return result, nil
}
