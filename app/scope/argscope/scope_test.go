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
	if value := scope.Value("$0"); value != "v1" {
		t.Errorf("$0 must be equal to position value %v != %v", value, "v1")
	}
	if value := scope.Value("$1"); value != "path=my/path" {
		t.Errorf("$1 must be equal to position value %v != %v", value, "path=my/path")
	}
	if value := scope.Value("$2"); value != "number=12" {
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
	var have = scope.Value(index)
	if have != expected {
		t.Errorf(`%s must be equal to "%v" and take "%v"`, index, expected, have)
	}
}
