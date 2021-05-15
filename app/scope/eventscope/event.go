package eventscope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
)

// EventScope is event scope interface
type EventScope struct {
	eventsCallbacks map[interface{}][]app.EventCallback
	mu              sync.RWMutex
}

// NewEventScope create new instance of event scope
func New() app.EventScope {
	return app.EventScope(&EventScope{
		eventsCallbacks: make(map[interface{}][]app.EventCallback),
	})
}

// Trigger run all function connected to event
func (es *EventScope) Trigger(eID interface{}, data interface{}) error {
	es.mu.RLock()
	defer es.mu.RUnlock()
	for _, onFunc := range es.eventsCallbacks[eID] {
		if err := onFunc(data); err != nil {
			return err
		}
	}
	return nil
}

// On connect a function to event
func (es *EventScope) On(eID interface{}, callback app.EventCallback) {
	es.mu.Lock()
	defer es.mu.Unlock()
	callbacks, ok := es.eventsCallbacks[eID]
	if !ok {
		es.eventsCallbacks[eID] = []app.EventCallback{callback}
		return
	}
	es.eventsCallbacks[eID] = append(callbacks, callback)
}
