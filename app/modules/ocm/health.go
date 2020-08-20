package ocm

import (
	"github.com/goatcms/goatcore/varutil/goaterr"

	"github.com/goatcms/goatcore/app"
)

// SandboxHealthChecker check if sandbox contains all dependencies
func SandboxHealthChecker(a app.App, ctxScope app.Scope) (msg string, err error) {
	if !hasPodman && !hasDocker {
		err = goaterr.Errorf("Docker or podman are required. Installed docker ( https://docs.docker.com/get-docker/ ) or podman ( https://podman.io/getting-started/installation )")
		return err.Error(), err
	}
	if hasDocker && !hasPodman {
		err = goaterr.Errorf("Docker is available. (Warning: podman is unavailable for oc.engime=podman)")
		return err.Error(), err
	}
	if !hasDocker && hasPodman {
		err = goaterr.Errorf("Podman is available. (Warning: docker is unavailable for oc.engime=docker)")
		return err.Error(), err
	}
	return "Docker and podman are available. You can set container system by oc.engine=docker|podman attribute). Default container system is podman (as more secure).", nil
}
