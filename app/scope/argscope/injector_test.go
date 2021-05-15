package argscope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope/datascope"
)

type TestDeps struct {
	FirstArgument string `argument:"$0"`
	Path          string `argument:"path"`
	Number        int    `argument:"number"`
	Flag          bool   `argument:"flag"`
	Optional      int    `argument:"?optional"`
}

func TestInjectArgsSimpleStory(t *testing.T) {
	t.Parallel()
	var (
		dataScope = datascope.New(make(map[interface{}]interface{}))
		err       error
	)
	if err = InjectArgs(dataScope, "name=v1"); err != nil {
		t.Error(err)
		return
	}
	if dataScope.Value("name").(string) != "v1" {
		t.Errorf("expected v1")
	}
}

func TestGet(t *testing.T) {
	t.Parallel()
	var (
		scp = datascope.New(make(map[interface{}]interface{}))
		err error
	)
	if err = InjectArgs(scp, []string{
		"firs",
		"path=my/path",
		"-number=12",
		"--some=true",
		"second",
	}...); err != nil {
		t.Error(err)
		return
	}
	checkTestValue(t, scp, "$0", "firs")
	checkTestValue(t, scp, "path", "my/path")
	checkTestValue(t, scp, "number", "12")
	checkTestValue(t, scp, "some", "true")
	checkTestValue(t, scp, "$1", "second")
}

func TestNewScopeFromString(t *testing.T) {
	t.Parallel()
	var scp = datascope.New(make(map[interface{}]interface{}))
	if err := InjectString(scp, `v1 path=my/path number="12 12" "--some=true true --" -- ignored=true`); err != nil {
		t.Error(err)
		return
	}
	checkTestValue(t, scp, "$0", "v1")
	checkTestValue(t, scp, "path", "my/path")
	checkTestValue(t, scp, "number", "12 12")
	checkTestValue(t, scp, "some", "true true --")
	checkTestValue(t, scp, "number", "12 12")
	checkTestValue(t, scp, "ignored", nil)
}

func checkTestValue(t *testing.T, scope app.DataScope, index string, expected interface{}) {
	if have := scope.Value(index); have != expected {
		t.Errorf(`%s must be equal to "%v" and take "%v"`, index, expected, have)
	}
}
