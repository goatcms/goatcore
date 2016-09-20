package scope

const (
	// KillEvent is kill action for current event
	KillEvent = 1
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
