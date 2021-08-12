package pool

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LockTest struct {
	suite.Suite
}

func (st *LockTest) SetupSuite() {

}

func (st *LockTest) TearDownSuite() {

}

func (st *LockTest) TestLockAndUnLock() {
	id := "testid"
	RLock(id)
	fmt.Println("one lock")

	RLock(id)
	fmt.Println("two lock")

	RUnlock(id)
	RUnlock(id)

	Lock(id)
	Unlock(id)
	// Unlock(id)
	// fmt.Println("wait")

}

func TestLockTest(t *testing.T) {
	suite.Run(t, new(LockTest))
}
