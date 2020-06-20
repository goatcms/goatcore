package goaterr

import (
	"errors"
	"strings"
	"testing"
)

func TestNewError(t *testing.T) {
	t.Parallel()
	var err = NewError("Some message")
	if err == nil {
		t.Errorf("Expected error (not nil)")
	}
	if err.Error() != "Some message" {
		t.Errorf("Result should contains 'Some message'")
	}
}

func TestToError(t *testing.T) {
	t.Parallel()
	var (
		err = ToError([]error{
			errors.New("First error"),
			errors.New("Second error"),
		})
		errJSON JSONError
		ok      bool
	)
	if err == nil {
		t.Errorf("Expected error (not nil)")
		return
	}
	if errJSON, ok = err.(JSONError); !ok {
		t.Errorf("Error should implement JSONError")
		return
	}
	json := errJSON.ErrorJSON()
	if !strings.Contains(json, "First error") {
		t.Errorf("Result should contains 'First error'")
	}
	if !strings.Contains(json, "Second error") {
		t.Errorf("Result should contains 'Second error'")
	}
}

func TestWrapToJSON(t *testing.T) {
	var (
		err     error
		errJSON JSONError
		ok      bool
	)
	t.Parallel()
	if err = Wrap(errors.New("Parent error"), "Child error"); err == nil {
		t.Errorf("Expected error (not nil)")
	}
	if errJSON, ok = err.(JSONError); !ok {
		t.Errorf("Error should implement JSONError")
		return
	}
	result := errJSON.ErrorJSON()
	if !strings.Contains(result, "Parent error") {
		t.Errorf("Result should contains 'Parent error' and take %s", result)
	}
	if !strings.Contains(result, "Child error") {
		t.Errorf("Result should contains 'Child error' and take %s", result)
	}
}
func TestWrapMessage(t *testing.T) {
	t.Parallel()
	var err = Wrap(errors.New("Parent error"), "Child error")
	if err == nil {
		t.Errorf("Expected error (not nil)")
	}
	if err.Error() == "Parent error" {
		t.Errorf("err.Error() should be equal to 'Parent error' and it is %s", err.Error())
	}
}
