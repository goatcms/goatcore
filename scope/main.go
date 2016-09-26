package scope

const (
	// KillEvent is kill action for current event
	KillEvent = 1
	// ErrorEvent is action for a error
	ErrorEvent = 1

	// Error is key for error value
	Error = "error"
)

// OnFunction is a roolback callback function
type OnFunction func(Scope) error

// Scope is global scope interface
type Scope interface {
	//DP() dependency.Provider

	Set(string, interface{})
	Get(string) interface{}

	Trigger(int) error
	On(int, OnFunction)
}
