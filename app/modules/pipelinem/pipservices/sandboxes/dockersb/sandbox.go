package dockersb

import (
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// DockerSandbox is termal sandbox
type DockerSandbox struct {
	imageName, cwd string
	deps           deps
}

// NewDockerSandbox create a DockerSandbox instance
func NewDockerSandbox(imageName string, deps deps) (ins pipservices.Sandbox, err error) {
	imageName = strings.Trim(imageName, " \t\n")
	if imageName == "" {
		return nil, goaterr.Errorf("Docker Sandbox: Container name can not be empty")
	}
	return &DockerSandbox{
		deps:      deps,
		imageName: imageName,
	}, nil
}

// Run run code in sandbox
func (sandbox *DockerSandbox) Run(ctx app.IOContext) (err error) {
	var (
		cio    = ctx.IO()
		ok     bool
		cwd    filesystem.LocalFilespace
		cwdAbs string
		envs   commservices.Environments
	)
	if cwd, ok = cio.CWD().(filesystem.LocalFilespace); !ok {
		return goaterr.Errorf("DockerSandbox support only filesystem.LocalFilespace as CWD (Current Working Directory) and take %T", cio.CWD())
	}
	if cwdAbs, err = filepath.Abs(cwd.LocalPath()); err != nil {
		return err
	}
	if envs, err = sandbox.deps.EnvironmentsUnit.Envs(ctx.Scope()); err != nil {
		return err
	}
	volumeAttr := `--volume=` + cwdAbs + `:/cwd`
	args := []string{"docker", "run", "-i", "--rm", "-w=/cwd", volumeAttr, "--entrypoint", "/bin/sh"}
	for key, value := range envs.GetAll() {
		args = append(args, "-e", key+"="+strconv.Quote(value))
	}
	args = append(args, sandbox.imageName)
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = cio.In()
	cmd.Stdout = cio.Out()
	cmd.Stderr = cio.Err()
	cmd.Dir = cwd.LocalPath()
	if err = cmd.Run(); err != nil {
		cio.Err().Printf(err.Error())
		err = goaterr.Wrapf("Docker sandbox error: %s", err, err.Error())
		return err
	}
	return nil
}
