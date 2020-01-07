package dockersb

import (
	"os/exec"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// DockerSandbox is termal sandbox
type DockerSandbox struct {
	imageName, cwd string
}

// NewDockerSandbox create a DockerSandbox instance
func NewDockerSandbox(imageName string) (ins pipservices.Sandbox, err error) {
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
		cio = ctx.IO()
		ok  bool
		cwd filesystem.LocalFilespace
	)
	if cwd, ok = cio.CWD().(filesystem.LocalFilespace); !ok {
		return goaterr.Errorf("DockerSandbox support only filesystem.LocalFilespace as CWD (Current Working Directory) and take %T", cio.CWD())
	}
	args := []string{"docker", "run", "-i", "--rm", "--entrypoint", "/bin/sh", sandbox.imageName}
	ctx.IO().Out().Printf("Run docker sandbox %s by %v", sandbox.imageName, args)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = cio.In()
	cmd.Stdout = cio.Out()
	cmd.Stderr = cio.Err()
	cmd.Dir = cwd.LocalPath()
	if err = cmd.Run(); err != nil {
		cio.Err().Printf("\nDocker sandbox error: %v\n", err)
		return goaterr.Wrap(err)
	}
	return nil
}
