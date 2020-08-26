package memfs

import (
	"testing"
)

func TestReducePathFails(t *testing.T) {
	t.Parallel()
	var err error
	if _, err = reducePath(".."); err == nil {
		t.Errorf("expected error for path '..'")
	}
	if _, err = reducePath("/a/../.."); err == nil {
		t.Errorf("expected error for path '/a/../..'")
	}
	if _, err = reducePath("./../.."); err == nil {
		t.Errorf("expected error for path './../..'")
	}
	if _, err = reducePath("./.."); err == nil {
		t.Errorf("expected error for path './..'")
	}
}

func TestReducePathSuccess(t *testing.T) {
	t.Parallel()
	var (
		err  error
		path string
	)
	if path, err = reducePath("./a"); err != nil {
		t.Error(err)
		return
	}
	if path != "a" {
		t.Errorf("expected a and take %s", path)
		return
	}
	if path, err = reducePath("a/b/c"); err != nil {
		t.Error(err)
		return
	}
	if path != "a/b/c" {
		t.Errorf("expected a/b/c and take %s", path)
		return
	}
	if path, err = reducePath("some/../result/path"); err != nil {
		t.Error(err)
		return
	}
	if path != "result/path" {
		t.Errorf("expected result/path and take %s", path)
		return
	}
}
