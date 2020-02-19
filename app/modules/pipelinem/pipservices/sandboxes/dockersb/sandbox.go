package dockersb

import (
	"io"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/app/modules/pipelinem/pipservices"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/varutil"
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
		cio        = ctx.IO()
		ok         bool
		cwd        filesystem.LocalFilespace
		cwdAbs     string
		envs       commservices.Environments
		initReader io.Reader
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
	args := []string{"docker", "run", "-i", "--rm", "-w=/cwd", volumeAttr, "--entrypoint", "/bin/sh", sandbox.imageName}
	if initReader, err = sandbox.initSequence(envs); err != nil {
		return err
	}
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = io.MultiReader(initReader, cio.In())
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

func (sandbox *DockerSandbox) initSequence(envs commservices.Environments) (reader io.Reader, err error) {
	var (
		initCode = "\nset -e\n"
		eofTag   = "EOF" + varutil.RandString(10, varutil.UpperAlphaBytes)
	)
	for key, value := range envs.All() {
		initCode += "export " + key + "=$(cat <<" + eofTag + "\n" + value + "\n" + eofTag + "\n)\n"
	}
	sshCert := envs.SSHCert()
	if sshCert.Public != "" || sshCert.Secret != "" {
		if sshCert.Public == "" {
			return nil, goaterr.Errorf("SSHCert: Public key is required")
		}
		if sshCert.Secret == "" {
			return nil, goaterr.Errorf("SSHCert: Secret key is required")
		}
		initCode += "mkdir -p ~/.ssh\n"
		if sshCert.Public != "" {
			initCode += "cat <<" + eofTag + " >> ~/.ssh/id_rsa.pub \n" + sshCert.Public + "\n" + eofTag + "\n"
		}
		if sshCert.Secret != "" {
			initCode += "cat <<" + eofTag + " >> ~/.ssh/id_rsa \n" + sshCert.Secret + "\n" + eofTag + "\n"
		}
	}
	return strings.NewReader(initCode), nil
}
