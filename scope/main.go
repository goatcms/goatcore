package scope

const (
	// KillEvent is kill action for current event
	KillEvent = iota
	// ErrorEvent is action for a error
	ErrorEvent = iota
	// CommitEvent is a action run when data is persist
	CommitEvent = iota
	// RollbackEvent is a action run when data is restore
	RollbackEvent = iota

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
