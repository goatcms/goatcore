package corescope

import (
	"fmt"
	"testing"

	"github.com/goatcms/goat-core/scope"
)

func TestSimpleStory(t *testing.T) {
	var called1 = false
	var called2 = false
	c := NewScope(nil)
	c.On(scope.KillEvent, func(s scope.Scope) error {
		called1 = true
		return nil
	})
	c.On(scope.KillEvent, func(s scope.Scope) error {
		called2 = true
		return nil
	})
	if err := c.Trigger(scope.KillEvent); err != nil {
		t.Error(err)
		return
	}
	if called1 == false || called2 == false {
		t.Errorf("Trigger don't run function connected to scope event")
	}
}

func TestErrorStory(t *testing.T) {
	c := NewScope(nil)
	c.On(scope.KillEvent, func(s scope.Scope) error {
		return fmt.Errorf("something is wrong")
	})
	if err := c.Trigger(scope.KillEvent); err == nil {
		t.Errorf("Trigger should return error if a function is failed")
	}
}
