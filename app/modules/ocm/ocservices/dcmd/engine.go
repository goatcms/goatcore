package dcmd

import (
	"io"
	"os/exec"
	"strings"

	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/app/modules/ocm/ocservices"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// Engine SandboxesEngine is a tool to menage sandboxes.
type Engine struct {
	appName string
}

// NewEngine create docekr engine instance
func NewEngine(appName string) ocservices.Engine {
	return &Engine{
		appName: appName,
	}
}

// Run container
func (engine *Engine) Run(container ocservices.Container) (err error) {
	var (
		cio = container.IO
		// ok         bool
		cwdAbs     string
		args       []string
		command    []string
		initReader io.Reader
	)
	// prepare command
	command = []string{
		engine.appName,
		"run",
		"-i",
		"--rm",
	}
	if container.Privileged {
		command = append(command, "--privileged")
	}
	// map volumes
	if args, err = MapVolumens(container.FSVolumes); err != nil {
		return err
	}
	command = append(command, args...)
	if container.WorkDir != "" {
		command = append(command, "-w="+container.WorkDir)
	}
	if container.Entrypoint != "" {
		command = append(command, "--entrypoint", container.Entrypoint)
	}
	command = append(command, container.Image)
	// prepare command
	if initReader, err = InitSequence(container.Envs); err != nil {
		return err
	}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdin = gio.NewSafeReader(io.MultiReader(initReader, cio.In()))
	cmd.Stdout = gio.NewSafeWriter(cio.Out())
	cmd.Stderr = gio.NewSafeWriter(cio.Err())
	cmd.Dir = cwdAbs
	if err = cmd.Run(); err != nil {
		err = goaterr.Wrapf("`%s` error: %s", err, strings.Join(command, " "), err.Error())
		cio.Err().Printf(err.Error())
		return err
	}
	return nil
}
