package argscope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
)

type TestDeps struct {
	FirstArgument string `argument:"$0"`
	Path          string `argument:"path"`
	Number        int    `argument:"number"`
	Flag          bool   `argument:"flag"`
	Optional      int    `argument:"?optional"`
}

/*
Important: We don't support types conversion any more. - 26.07.2019
func TestInjectTo(t *testing.T) {
	t.Parallel()
	var deps TestDeps
	scope, err := NewScope([]string{"v1", "path=my/path", "number=12", "flag=1"}, app.ArgsTagName)
	if err != nil {
		t.Error(err)
		return
	}
	if err = scope.InjectTo(&deps); err != nil {
		t.Error(err)
		return
	}
	if deps.FirstArgument != "v1" {
		t.Errorf("FirstArgument must be equal to first value %v != %v", deps.FirstArgument, "v1")
	}
	if deps.Path != "my/path" {
		t.Errorf("Path must be equal to path argument %v != %v", deps.Path, "my/path")
	}
	if deps.Number != 12 {
		t.Errorf("FirstArgument must be equal to first value %v != %v", deps.Number, 12)
	}
	if deps.Flag != true {
		t.Errorf("Flag must be equal to first true %v != %v", deps.Flag, true)
	}
}
*/

func TestInjectToFail(t *testing.T) {
	t.Parallel()
	var deps TestDeps
	scope, err := NewScope([]string{"v1"}, app.ArgsTagName)
	if err != nil {
		t.Error(err)
		return
	}
	if err := scope.InjectTo(&deps); err == nil {
		t.Errorf("Should return error when all required deps are defined")
		return
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	scope, err := NewScope([]string{"v1", "path=my/path", "number=12", "--some=true"}, app.ArgsTagName)
	if err != nil {
		t.Error(err)
		return
	}
	if value, _ := scope.Get("$0"); value != "v1" {
		t.Errorf("$0 must be equal to position value %v != %v", value, "v1")
	}
	if value, _ := scope.Get("$1"); value != "path=my/path" {
		t.Errorf("$1 must be equal to position value %v != %v", value, "path=my/path")
	}
	if value, _ := scope.Get("$2"); value != "number=12" {
		t.Errorf("$2 must be equal to position value %v != %v", value, "number=12")
	}
}

func TestNewScopeFromString(t *testing.T) {
	t.Parallel()
	scope, err := NewScopeFromString(`v1 path=my/path number="12 12" "--some=true true"`, app.ArgsTagName)
	if err != nil {
		t.Error(err)
		return
	}
	checkTestValue(t, scope, "$0", "v1")
	checkTestValue(t, scope, "$1", "path=my/path")
	checkTestValue(t, scope, "$2", "number=12 12")
	checkTestValue(t, scope, "$3", "--some=true true")
	checkTestValue(t, scope, "number", "12 12")
}

func checkTestValue(t *testing.T, scope app.Scope, index, expected string) {
	var (
		have interface{}
		err  error
	)
	if have, err = scope.Get(index); err != nil {
		t.Error(err)
		return
	}
	if have != expected {
		t.Errorf(`%s must be equal to "%v" and take "%v"`, index, expected, have)
	}
}
