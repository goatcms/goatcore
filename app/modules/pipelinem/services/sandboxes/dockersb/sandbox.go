package dockersb

import (
	"os/exec"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/services"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// DockerSandbox is termal sandbox
type DockerSandbox struct {
	imageName, cwd string
}

// NewDockerSandbox create a DockerSandbox instance
func NewDockerSandbox(imageName string) (ins services.Sandbox, err error) {
	imageName = strings.Trim(imageName, " \t\n")
	if imageName == "" {
		return nil, goaterr.Errorf("Docker Sandbox: Container name can not be empty")
	}
	return &DockerSandbox{
		imageName: imageName,
	}, nil
}

// Run run code in sandbox
func (sandbox *DockerSandbox) Run(ctx app.IOContext) (err error) {
	var (
		io  = ctx.IO()
		ok  bool
		cwd filesystem.LocalFilespace
	)
	if cwd, ok = io.CWD().(filesystem.LocalFilespace); !ok {
		return goaterr.Errorf("DockerSandbox support only filesystem.LocalFilespace as CWD (Current Working Directory) and take %T", io.CWD())
	}
	cmd := exec.Command("docker", "run", "-it", "--rm", "--entrypoint", "/bin/sh", sandbox.imageName)
	cmd.Stdin = io.In()
	cmd.Stdout = io.Out()
	cmd.Stderr = io.Err()
	cmd.Dir = cwd.LocalPath()
	err = cmd.Run()
	return goaterr.Wrap(err)
}
