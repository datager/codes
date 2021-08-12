package clients

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LinkClientTest struct {
	suite.Suite
	linkClient *LinkClient
}

func (st *LinkClientTest) SetupSuite() {
	SetupTestEnv()
	testService := GetTestService()

	st.linkClient = testService.linkClient
}

func (st *LinkClientTest) TeardownSuite() {
	TeardownTestEnv()
}

func (st *LinkClientTest) TestLink() {
	linkType := LinkTypeImporterLibraf
	worker := 1
	engineType := EngineTypeVehicle

	dir, err := os.Getwd()
	st.Equal(nil, err)
	r := strings.Split(dir, "loki")

	path := r[0] + "/loki/config/templates/importer_libraF.json"

	content, err := ioutil.ReadFile(path)
	st.Equal(nil, err)

	// delete space
	str := strings.Replace(string(content), " ", "", -1)

	// delete \n \r
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)

	result, err := st.linkClient.CreateLinkTask(linkType, worker, engineType, str)
	st.Equal(nil, err)
	st.Equal(linkType, result.Type)
	st.Equal(worker, result.Workers)
	st.Equal(engineType, result.Engine)
	// st.Equal(RUNNING, result.State)

	result, err = st.linkClient.GetLinkTaskByID(result.InstanceID)
	st.Equal(nil, err)
	st.Equal(linkType, result.Type)
	// st.Equal(RUNNING, result.State)

	list, err := st.linkClient.GetLinkTasks()
	st.Equal(nil, err)
	if len(list) == 1 {
		st.Equal(result.InstanceID, list[0].InstanceID)
	}

	result, err = st.linkClient.StopLinkTask(result.InstanceID)
	st.Equal(nil, err)
	// st.Equal(STOPPED, result.State)

	result, err = st.linkClient.StartLinkTask(result.InstanceID)
	st.Equal(nil, err)
	// st.Equal(RUNNING, result.State)

	limitData, err := st.linkClient.GetNodeLimit()
	st.Equal(nil, err)
	fmt.Println(limitData)

	result, err = st.linkClient.DeleteLinkTask(result.InstanceID)
	st.Equal(nil, err)
	st.Equal("DELETED", result.State)

}

func TestLinkClient(t *testing.T) {
	suite.Run(t, new(LinkClientTest))
}
