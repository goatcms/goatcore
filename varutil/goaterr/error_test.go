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
	result := err.Error()
	if !strings.Contains(result, "Some message") {
		t.Errorf("Result should contains 'Some message'")
	}
}

func TestToError(t *testing.T) {
	t.Parallel()
	var err = ToError([]error{
		errors.New("First error"),
		errors.New("Second error"),
	})
	if err == nil {
		t.Errorf("Expected error (not nil)")
	}
	result := err.Error()
	if !strings.Contains(result, "First error") {
		t.Errorf("Result should contains 'First error'")
	}
	if !strings.Contains(result, "Second error") {
		t.Errorf("Result should contains 'Second error'")
	}
}

func TestWrapError(t *testing.T) {
	t.Parallel()
	var err = Wrap(errors.New("Parent error"), "Child error")
	if err == nil {
		t.Errorf("Expected error (not nil)")
	}
	result := err.Error()
	if !strings.Contains(result, "Parent error") {
		t.Errorf("Result should contains 'Parent error' and take %s", result)
	}
	if !strings.Contains(result, "Child error") {
		t.Errorf("Result should contains 'Child error' and take %s", result)
	}
}
