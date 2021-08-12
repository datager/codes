package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestQueue struct {
	suite.Suite
}

func (st *TestQueue) SetupSuite() {

}

func (st *TestQueue) TeardownSuite() {

}

func (st *TestQueue) Test() {
	q := Queue{}

	index := q.EnQueue(1)
	st.Equal(0, index)

	index = q.EnQueue(2)
	st.Equal(1, index)

	v := q.Next()
	st.Equal(1, v)

	v = q.Next()
	st.Equal(2, v)

	v = q.Next()
	st.Equal(nil, v)

	err := q.DeQueueWithIndex(0)
	st.Equal(nil, err)

	v = q.Next()
	st.Equal(2, v)

}

func TestQueueRun(t *testing.T) {
	suite.Run(t, new(TestQueue))
}
