package scope

/*
// ChildScope represent sub-scope
type ChildScope struct {
	parent   app.Scope
	data     DataScope
	events   EventScope
	injector app.Injector
}

// NewChildScope create new instance of scope
func NewChildScope(parent app.Scope, tagname string) app.Scope {
	ds := DataScope{}
	return &ChildScope{
		parent:   parent,
		data:     ds,
		injector: ds.Injector(tagname),
		events:   EventScope{},
	}
}

// Set new scope value
func (cs *ChildScope) Set(key string, v interface{}) error {
	return cs.data.Set(key, v)
}

// Get get value from context
func (cs *ChildScope) Get(key string) interface{} {
	val := cs.data.Get(key)
	if val != nil {
		return val
	}
	return cs.parent.Get(key)
}

// Trigger run all function connected to event
func (cs *ChildScope) Trigger(eID int) error {
	if err := cs.events.Trigger(eID); err != nil {
		return err
	}
	if err := cs.parent.Trigger(eID); err != nil {
		return err
	}
	return nil
}

// On connect a function to event
func (cs *ChildScope) On(eID int, callback app.Callback) {
	cs.events.On(eID, callback)
}

// InjectTo insert data to object
func (cs *ChildScope) InjectTo(obj interface{}) error {
	if err := cs.injector.InjectTo(obj); err != nil {
		return err
	}
	if err := cs.parent.InjectTo(obj); err != nil {
		return err
	}
	return nil
}
*/
