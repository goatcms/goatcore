package scope

import (
	"sync"

	"github.com/goatcms/goatcore/app"
)

// ChildEventScope is event scope interface
type ChildEventScope struct {
	parent    app.EventScope
	callbacks map[int][]app.EventCallback
	mu        sync.RWMutex
}

// NewChildEventScope create new instance of event scope
func NewChildEventScope(parent app.EventScope) app.EventScope {
	return app.EventScope(&ChildEventScope{
		parent:    parent,
		callbacks: make(map[int][]app.EventCallback),
	})
}

// Trigger run all function connected to event
func (es *ChildEventScope) Trigger(eID int, data interface{}) (err error) {
	if err = es.parent.Trigger(eID, data); err != nil {
		return err
	}
	es.mu.RLock()
	defer es.mu.RUnlock()
	for _, onFunc := range es.callbacks[eID] {
		if err := onFunc(data); err != nil {
			return err
		}
	}
	return nil
}

// On connect a function to event
func (es *ChildEventScope) On(eID int, callback app.EventCallback) {
	es.mu.Lock()
	defer es.mu.Unlock()
	callbacks, ok := es.callbacks[eID]
	if !ok {
		es.callbacks[eID] = []app.EventCallback{callback}
		return
	}
	es.callbacks[eID] = append(callbacks, callback)
}
