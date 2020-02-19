package varg

import "testing"

func TestMatchBoolForValue(t *testing.T) {
	t.Parallel()
	var (
		result bool
		err    error
	)
	if result, err = MatchBool("name", "true", false); err != nil {
		t.Error(err)
		return
	}
	if result != true {
		t.Errorf("Expected true and take '%v'", result)
	}
}

func TestMatchBoolDefaultValue(t *testing.T) {
	t.Parallel()
	var (
		result bool
		err    error
	)
	if result, err = MatchBool("name", "", false); err != nil {
		t.Error(err)
		return
	}
	if result != false {
		t.Errorf("Expected false and take '%v'", result)
	}
}

func TestMatchBoolIncorrectValue(t *testing.T) {
	t.Parallel()
	var (
		err error
	)
	if _, err = MatchBool("name", "wrong-value", false); err == nil {
		t.Errorf("Expected error")
		return
	}
}
