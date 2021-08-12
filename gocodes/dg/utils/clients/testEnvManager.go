package clients

import (
	"sync"
)

type TestEnvManager struct {
	IsTestEnvInitialed bool
	TestCounter        int
}

var initialLock sync.Mutex
var finalLock sync.Mutex

var testEnvMng = &TestEnvManager{IsTestEnvInitialed: false}

var testService *TestService

func SetupTestEnv() {
	initialLock.Lock()
	defer initialLock.Unlock()
	if !testEnvMng.IsTestEnvInitialed {
		testEnvMng.IsTestEnvInitialed = true

		testService = InitTestEnv()

	}
	testEnvMng.TestCounter = testEnvMng.TestCounter + 1
}

func TeardownTestEnv() {
	finalLock.Lock()
	defer finalLock.Unlock()

	testEnvMng.TestCounter = testEnvMng.TestCounter - 1
	if testEnvMng.TestCounter == 0 {
		CleanTestEnv()
		testEnvMng.IsTestEnvInitialed = false
	}
}

func GetTestService() *TestService {
	return testService
}
