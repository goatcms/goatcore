package dcmd

import (
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil"

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
		cio        = container.IO
		cwdAbs     string
		args       []string
		command    []string
		initReader io.Reader
		cmd        *exec.Cmd
		name       string
	)
	// prepare name
	cTime := time.Now()
	name = "dcmd" + cTime.Format("2006_01_02_15_04_05") + varutil.RandString(7, varutil.AlphaNumericBytes)
	// prepare command
	command = []string{
		engine.appName,
		"run",
		"-i",
		"--rm",
		"--name",
		name,
	}
	if container.Privileged {
		command = append(command, "--privileged")
	}
	if args, err = MapVolumens(container.FSVolumes); err != nil {
		return err
	}
	command = append(command, args...)
	if args, err = MapPorts(container.Ports); err != nil {
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
	if container.Scope == nil {
		cmd = exec.Command(command[0], command[1:]...)
	} else {
		cmd = exec.CommandContext(container.Scope.Context(), command[0], command[1:]...)
		if err = container.Scope.AddTasks(1); err != nil {
			return err
		}
		defer container.Scope.DoneTask()
		container.Scope.On(app.KillEvent, func(interface{}) error {
			return killContainer(name)
		})
	}
	cmd.Stdin = gio.NewSafeReader(io.MultiReader(initReader, cio.In()))
	cmd.Stdout = gio.NewSafeWriter(cio.Out())
	cmd.Stderr = gio.NewSafeWriter(cio.Err())
	cmd.Dir = cwdAbs
	if err = cmd.Run(); err != nil {
		err = goaterr.Wrapf("`%s` error: %s", err, strings.Join(command, " "), err.Error())
		cio.Err().Printf(err.Error())
		if container.Scope != nil {
			container.Scope.AppendError(err)
		}
		return err
	}
	return nil
}

func killContainer(name string) (err error) {
	killCmd := exec.Command("docker", "kill", name)
	return killCmd.Run()
}
