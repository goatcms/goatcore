package scope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

func TestEventStory(t *testing.T) {
	t.Parallel()
	var called1 = false
	var called2 = false
	c := NewEventScope()
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

func TestErrorStory(t *testing.T) {
	t.Parallel()
	c := NewEventScope()
	c.On(app.KillEvent, func(interface{}) error {
		return goaterr.Errorf("something is wrong")
	})
	if err := c.Trigger(app.KillEvent, nil); err == nil {
		t.Errorf("Trigger should return error if a function is failed")
	}
}
