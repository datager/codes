package clients

import (
	"fmt"
	"codes/gocodes/dg/utils"
	"codes/gocodes/dg/utils/config"
)

type TestService struct {
	mserverClient *MServerClient
	linkClient    *LinkClient
	vseMatrix     *VseMatrixHTTPComposed
	athenaClient  *AthenaClient
}

func InitTestEnv() *TestService {
	// os.Setenv("ENVIRONMENT", "test")
	conf, err := config.InitConfig("../../config")
	if err != nil {
		panic(err)
	}

	mserverAddr := conf.GetString("services.mserver.run_addr")
	timeout := utils.CLIENT_TIMEOUT
	outAddr := conf.GetString("services.mserver.out_addr")

	mserverClient := NewMServerClient(mserverAddr, timeout, outAddr)

	linkAddr := conf.GetString("services.link.addr")

	linkClient := NewLinkClient(linkAddr, timeout)

	athenaAddr := conf.GetString("services.athena.addr")
	athenaClient := NewAthenaClient(athenaAddr, timeout)

	faceAddr := conf.GetString(fmt.Sprintf("services.%v.addr", "vse_face"))
	vehicleAddr := conf.GetString(fmt.Sprintf("services.%v.addr", "vse_vehicle"))
	// allAddr := conf.GetString(fmt.Sprintf("services.%v.addr", "vse_allobj"))

	timeoutFace := conf.GetDurationOrDefault(fmt.Sprintf("services.%v.timeout", "vse_face"), utils.CLIENT_TIMEOUT)
	// number := conf.GetIntOrDefault(fmt.Sprintf("services.%v.connect_number", "vse_face"), 3)
	// tryTimes := conf.GetIntOrDefault(fmt.Sprintf("services.%v.try_times", "vse_face"), 3)

	faceVse := NewVseMatrixHTTPClient(faceAddr, timeoutFace)
	vehicleVse := NewVseMatrixHTTPClient(vehicleAddr, timeoutFace)
	// allVse := NewVseMatrixPoolService(allAddr, number, timeoutFace, tryTimes)

	param := VseMatrixHTTPComposedParams{
		VseFaceClient:    faceVse,
		VseVehicleClient: vehicleVse,
		// VseAllObjClient:  allVse,
	}

	vseMatrx, err := NewVseMatrixHTTPComposed(param)
	if err != nil {
		panic(err)
	}

	return &TestService{
		mserverClient: mserverClient,
		linkClient:    linkClient,
		vseMatrix:     vseMatrx,
		athenaClient:  athenaClient,
	}

}

func CleanTestEnv() {

}
