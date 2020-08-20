package containersb

import (
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// ContainerSandbox is termal sandbox
type ContainerSandbox struct {
	imageName  string
	cwd        string
	entrypoint string
	deps       deps
}

// NewContainerSandbox create a ContainerSandbox instance
func NewContainerSandbox(imageName, entrypoint string, deps deps) (ins pipservices.Sandbox, err error) {
	imageName = strings.Trim(imageName, " \t\n")
	if imageName == "" {
		return nil, goaterr.Errorf("Docker Sandbox: Container name can not be empty")
	}
	return &ContainerSandbox{
		deps:       deps,
		imageName:  imageName,
		entrypoint: entrypoint,
	}, nil
}

// Run run code in sandbox
func (sandbox *ContainerSandbox) Run(ctx app.IOContext) (err error) {
	var (
		envs commservices.Environments
	)
	if envs, err = sandbox.deps.EnvironmentsUnit.Envs(ctx.Scope()); err != nil {
		return err
	}
	return sandbox.deps.OCManager.Run(ocservices.Container{
		IO:         ctx.IO(),
		Image:      sandbox.imageName,
		WorkDir:    "/cwd",
		Entrypoint: sandbox.entrypoint,
		Envs:       envs,
		FSVolumes: map[string]ocservices.FSVolume{
			"/cwd": ocservices.FSVolume{
				Filespace: ctx.IO().CWD(),
			},
		},
		Provilages: false,
	})
}
