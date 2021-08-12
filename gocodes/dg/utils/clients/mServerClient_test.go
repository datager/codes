package clients

import (
	"codes/gocodes/dg/models"
	"codes/gocodes/dg/utils/log"

	"github.com/stretchr/testify/suite"
)

type MServerClientTest struct {
	suite.Suite
	mServerClient *MServerClient
}

func (st *MServerClientTest) SetupSuite() {
	SetupTestEnv()
	testService := GetTestService()

	st.mServerClient = testService.mserverClient
}

func (st *MServerClientTest) TearDownSuite() {
	TeardownTestEnv()
}

func (st *MServerClientTest) TestMServerClient() {
	sensorType := models.SensorType_Sensor_Type_Ipc
	sensorID := "testsensor"
	req := RunRequest{
		Input:        "rtsp://192.168.2.123/live/t1065",
		Output:       "dmi://192.168.2.222:8905/live/" + sensorID,
		Args:         "",
		Type:         sensorType,
		Gb28181Param: models.GB28181Param{},
	}

	_, err := st.mServerClient.StartStream(req)
	st.Equal(nil, err)

	exist, err := st.mServerClient.CheckExist(sensorID)
	st.Equal(nil, err)
	st.Equal(true, exist)

	steam, err := st.mServerClient.GetStreamByID(sensorID)
	st.Equal(nil, err)
	st.Equal(sensorID, steam.StreamID)
	log.Debugln(*steam)

	err = st.mServerClient.StopStream(sensorID)
	st.Equal(nil, err)

	should := st.mServerClient.ShouldRegister(sensorType, "")
	st.Equal(true, should)
}

// func TestMServerClient(t *testing.T) {
// 	suite.Run(t, new(MServerClientTest))
// }
