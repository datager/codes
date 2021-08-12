package models

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestPreProcessSuite(t *testing.T) {
	suite.Run(t, new(PreProcessTest))
}

type PreProcessTest struct {
	suite.Suite
	req FaceConditionRequest

	testFuncs []func()
}

func (tst *PreProcessTest) SetupSuite() {
	tst.req = FaceConditionRequest{
		Vid: "8062c929-77d8-4402-9248-cd82ac690bd2",
		Pagination: Pagination{
			Limit:  0,
			Offset: 0,
		},
		OrderBy: "ts",
		Time: StartAndEndTimestamp{
			StartTimestamp: 1585031485359,
			EndTimestamp:   1585117885359,
		},
	}
}

func (tst *PreProcessTest) TearDownSuite() {

}

func (tst *PreProcessTest) TestRun() {
	tst.req.OrderAsc, tst.req.FromTailPage = false, false
	tst.print()

	tst.req.OrderAsc, tst.req.FromTailPage = false, true
	tst.print()

	tst.req.OrderAsc, tst.req.FromTailPage = true, false
	tst.print()

	tst.req.OrderAsc, tst.req.FromTailPage = true, true
	tst.print()
}

func (tst *PreProcessTest) print() {
	tst.req.PreProcess()
	if tst.req.OrderAsc == false && tst.req.FromTailPage == false {
		tst.T().Log("首页 降序 => 因为来自首页 所以正常按OrderAsc查, 不需再处理")
	} else if tst.req.OrderAsc == false && tst.req.FromTailPage == true {
		tst.T().Log("尾页 降序 => 因为来自尾页 所以按!orderAsc查, 然后在内存reverse")
	} else if tst.req.OrderAsc == true && tst.req.FromTailPage == false {
		tst.T().Log("首页 升序 => 因为来自首页, 所以正常按OrderAsc查, 不需再处理")
	} else if tst.req.OrderAsc == true && tst.req.FromTailPage == true {
		tst.T().Log("尾页 升序 => 因为来自尾页, 所以按!orderAsc查, 然后在内存reverse")
	} else {
	}
	tst.T().Log(tst.req.AscWithFromPage, tst.req.TryLimit)
}
