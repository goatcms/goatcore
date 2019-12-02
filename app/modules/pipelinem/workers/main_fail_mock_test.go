package workers

import (
	"fmt"

	"github.com/goatcms/webslots/slotsapp/services"
)

// MockedFailScriptExecutor is a mock for failed script executor
type MockedFailScriptExecutor struct {
	runned bool
}

// NewMockedScriptExecutor create new mock script executor
func NewMockedFailScriptExecutor() *MockedFailScriptExecutor {
	return &MockedFailScriptExecutor{
		runned: false,
	}
}

// Run script
func (mockExecutor *MockedFailScriptExecutor) Run(response services.ResponseContext) (err error) {
	mockExecutor.runned = true
	return fmt.Errorf("some fail")
}

func (mockExecutor *MockedFailScriptExecutor) Runned() bool {
	return mockExecutor.runned
}
