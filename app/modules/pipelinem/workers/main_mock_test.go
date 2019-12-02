package workers

import "github.com/goatcms/webslots/slotsapp/services"

// MockedScriptExecutor is a mock for script executor
type MockedScriptExecutor struct {
	runned bool
}

// NewMockedScriptExecutor create new mock script executor
func NewMockedScriptExecutor() *MockedScriptExecutor {
	return &MockedScriptExecutor{
		runned: false,
	}
}

// Run script
func (mockExecutor *MockedScriptExecutor) Run(response services.ResponseContext) (err error) {
	mockExecutor.runned = true
	return nil
}

func (mockExecutor *MockedScriptExecutor) Runned() bool {
	return mockExecutor.runned
}
