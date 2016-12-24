package scope

import "github.com/goatcms/goat-core/app"

// EventScope is event scope interface
type EventScope struct {
	eventsCallbacks map[int][]app.EventCallback
}

// NewEventScope create new instance of event scope
func NewEventScope() app.EventScope {
	return app.EventScope(&EventScope{
		eventsCallbacks: make(map[int][]app.EventCallback),
	})
}

// Trigger run all function connected to event
func (es *EventScope) Trigger(eID int, data interface{}) error {
	for _, onFunc := range es.eventsCallbacks[eID] {
		if err := onFunc(data); err != nil {
			return err
		}
	}
	return nil
}

// On connect a function to event
func (es *EventScope) On(eID int, callback app.EventCallback) {
	callbacks, ok := es.eventsCallbacks[eID]
	if !ok {
		es.eventsCallbacks[eID] = []app.EventCallback{callback}
		return
	}
	es.eventsCallbacks[eID] = append(callbacks, callback)
}
