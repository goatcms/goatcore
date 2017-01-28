package argscope

import (
	"testing"

	"github.com/goatcms/goat-core/app"
)

type TestDeps struct {
	FirstArgument string `argument:"$0"`
	Path          string `argument:"path"`
	Number        int    `argument:"number"`
	Flag          bool   `argument:"flag"`
	Optional      int    `argument:"?optional"`
}

func TestInjectTo(t *testing.T) {
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

func TestInjectToFail(t *testing.T) {
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
	scope, err := NewScope([]string{"v1", "path=my/path", "number=12", "--some=true"}, app.ArgsTagName)
	if err != nil {
		t.Error(err)
		return
	}
	if value, _ := scope.Get("$0"); value != "v1" {
		t.Errorf("$0 must be equal to position value %v != %v", value, "v1")
	}
	if value, _ := scope.Get("$0"); value != "v1" {
		t.Errorf("$1 must be equal to position value %v != %v", value, "path=my/path")
	}
	if value, _ := scope.Get("$0"); value != "v1" {
		t.Errorf("$2 must be equal to position value %v != %v", value, "number=12")
	}
}
