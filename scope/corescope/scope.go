package corescope

import (
	"github.com/goatcms/goat-core/dependency"
	"github.com/goatcms/goat-core/scope"
)

// Scope is global scope interface
type Scope struct {
	dp              dependency.Provider
	eventsCallbacks map[int][]scope.OnFunction
	values          map[string]interface{}
}

// NewScope create new instance of scope
func NewScope(dp dependency.Provider) scope.Scope {
	return &Scope{
		dp:              dp,
		eventsCallbacks: make(map[int][]scope.OnFunction),
		values:          make(map[string]interface{}),
	}
}

// DP returen scope dependency provider
func (s *Scope) DP() dependency.Provider {
	return s.dp
}

// Set new scope value
func (s *Scope) Set(key string, v interface{}) {
	s.values[key] = v
}

// Get get value from context
func (s *Scope) Get(key string) interface{} {
	return s.values[key]
}

// Trigger run all function connected to event
func (s *Scope) Trigger(eID int) error {
	for _, onFunc := range s.eventsCallbacks[eID] {
		if err := onFunc(s); err != nil {
			return err
		}
	}
	return nil
}

// On connect a function to event
func (s *Scope) On(eID int, callback scope.OnFunction) {
	callbacks, ok := s.eventsCallbacks[eID]
	if !ok {
		s.eventsCallbacks[eID] = []scope.OnFunction{callback}
		return
	}
	s.eventsCallbacks[eID] = append(callbacks, callback)
}
