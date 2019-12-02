package pipelinem

import (
	"os/exec"

	"github.com/goatcms/goatcore/app"
)

// SandboxHealthChecker check if sandbox contains all dependencies
func SandboxHealthChecker(a app.App, ctxScope app.Scope) (msg string, err error) {
	if err = exec.Command("docker", "version").Run(); err == nil {
		return "Workers docker sandbox require docker (install: https://www.docker.com )", err
	}
	return "Terminal sanbox", nil
}
