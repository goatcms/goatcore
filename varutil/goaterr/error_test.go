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
	if err = Wrap("Error wrapper", errors.New("Wraped error")); err == nil {
		t.Errorf("Expected error (not nil)")
	}
	if errJSON, ok = err.(JSONError); !ok {
		t.Errorf("Error should implement JSONError")
		return
	}
	result := errJSON.ErrorJSON()
	if !strings.Contains(result, "Error wrapper") {
		t.Errorf("Result should contains 'Parent error' and take %s", result)
	}
	if !strings.Contains(result, "Wraped error") {
		t.Errorf("Result should contains 'Wraped error' and take %s", result)
	}
}
func TestWrapMessage(t *testing.T) {
	t.Parallel()
	var err = Wrap("Error wrapper", errors.New("Wraped error"))
	if err == nil {
		t.Errorf("Expected error (not nil)")
	}
	if err.Error() == "Error wrapper" {
		t.Errorf("err.Error() should be equal to 'Error wrapper' and it is %s", err.Error())
	}
}
