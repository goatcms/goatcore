package dcmd

import (
	"testing"

	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"

	"github.com/goatcms/goatcore/testbase"
)

func TestDockerEngineStory(t *testing.T) {
	var (
		engine ocservices.Engine
		err    error
	)
	if _, err = testbase.LoadDockerTestConfig(); err != nil {
		t.Skip(err.Error())
		return
	}
	engine = NewEngine("docker")
	engineSimpleStory(t, engine)
}

func TestDockerPortStory(t *testing.T) {
	var (
		engine ocservices.Engine
		err    error
	)
	if _, err = testbase.LoadDockerTestConfig(); err != nil {
		t.Skip(err.Error())
		return
	}
	engine = NewEngine("docker")
	enginePortStory(t, engine)
}
