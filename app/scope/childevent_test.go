package scope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestChildEventScopeDoneStory(t *testing.T) {
	t.Parallel()
	var called1 = false
	var called2 = false
	c := NewChildEventScope(NewEventScope())
	c.On(app.KillEvent, func(interface{}) error {
		called1 = true
		return nil
	})
	c.On(app.KillEvent, func(interface{}) error {
		called2 = true
		return nil
	})
	if err := c.Trigger(app.KillEvent, nil); err != nil {
		t.Error(err)
		return
	}
	if called1 == false || called2 == false {
		t.Errorf("Trigger don't run function connected to scope event")
	}
}

func TestChildEventScopeParentDoneStory(t *testing.T) {
	t.Parallel()
	var called1 = false
	var called2 = false
	parentScp := NewEventScope()
	childScp := NewChildEventScope(parentScp)
	parentScp.On(app.KillEvent, func(interface{}) error {
		called1 = true
		return nil
	})
	parentScp.On(app.KillEvent, func(interface{}) error {
		called2 = true
		return nil
	})
	if err := childScp.Trigger(app.KillEvent, nil); err != nil {
		t.Error(err)
		return
	}
	if called1 == false || called2 == false {
		t.Errorf("Trigger don't run function connected to scope event")
	}
}

func TestChildEventScopeErrorStory(t *testing.T) {
	t.Parallel()
	c := NewChildEventScope(NewEventScope())
	c.On(app.KillEvent, func(interface{}) error {
		return goaterr.Errorf("something is wrong")
	})
	if err := c.Trigger(app.KillEvent, nil); err == nil {
		t.Errorf("Trigger should return error if a function is failed")
	}
}

func TestChildEventScopeParentErrorStory(t *testing.T) {
	t.Parallel()
	parentScp := NewEventScope()
	childScp := NewChildEventScope(parentScp)
	parentScp.On(app.KillEvent, func(interface{}) error {
		return goaterr.Errorf("something is wrong")
	})
	if err := childScp.Trigger(app.KillEvent, nil); err == nil {
		t.Errorf("Trigger should return error if a function is failed")
	}
}
