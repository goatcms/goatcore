package ocm

import (
	"os/exec"
)

var (
	hasDocker bool = exec.Command("docker", "version").Run() == nil
	hasPodman bool = exec.Command("podman", "version").Run() == nil
)
