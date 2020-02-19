package varg

import "testing"

func TestMatchStringForValue(t *testing.T) {
	t.Parallel()
	var (
		result string
		err    error
	)
	if result, err = MatchString("name", "VALUE1", []string{"value1"}, ""); err != nil {
		t.Error(err)
		return
	}
	if result != "value1" {
		t.Errorf("Expected 'value1' and take '%s'", result)
	}
}

func TestMatchStringDefault(t *testing.T) {
	t.Parallel()
	var (
		result string
		err    error
	)
	if result, err = MatchString("name", "", []string{"value1"}, "default"); err != nil {
		t.Error(err)
		return
	}
	if result != "default" {
		t.Errorf("Expected 'default' and take '%s'", result)
	}
}

func TestMatchStringIncorrect(t *testing.T) {
	t.Parallel()
	var (
		err error
	)
	if _, err = MatchString("name", "otherValue", []string{"value1"}, "default"); err == nil {
		t.Errorf("Expected error")
		return
	}
}
