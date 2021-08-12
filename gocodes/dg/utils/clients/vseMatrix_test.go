package clients

import (
	"github.com/stretchr/testify/suite"
)

type VseMatrixClientTest struct {
	suite.Suite
	client *VseMatrixHTTPComposed
}

func (st *VseMatrixClientTest) SetupSuite() {
	SetupTestEnv()
	testService := GetTestService()

	st.client = testService.vseMatrix
}

func (st *VseMatrixClientTest) TearDownSuite() {
	TeardownTestEnv()
}

//func (st *VseMatrixClientTest) TestGetIndex() {
//	r, err := st.client.GetIndexFromVehicle(dg_model.IndexType_INDEX_FACE)
//	fmt.Println("index:", r)
//	st.Equal(nil, err)
//}
//
//func TestVse(t *testing.T) {
//	suite.Run(t, new(VseMatrixClientTest))
//}
