package dcmd

import (
	"testing"

	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"

	"github.com/goatcms/goatcore/testbase"
)

func TestPodmanEngineStory(t *testing.T) {
	var (
		engine ocservices.Engine
		err    error
	)
	if _, err = testbase.LoadPodmanTestConfig(); err != nil {
		t.Skip(err.Error())
		return
	}
	engine = NewEngine("podman")
	engineSimpleStory(t, engine)
}

func TestPodmanPortStory(t *testing.T) {
	var (
		engine ocservices.Engine
		err    error
	)
	if _, err = testbase.LoadPodmanTestConfig(); err != nil {
		t.Skip(err.Error())
		return
	}
	engine = NewEngine("podman")
	enginePortStory(t, engine)
}
