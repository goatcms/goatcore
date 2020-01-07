package testbase

import (
	"fmt"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// DockerTestConfig is docker test config object
type DockerTestConfig struct {
	onStr string
}

// LoadDockerTestConfig return docker config from envs
func LoadDockerTestConfig() (config *DockerTestConfig, err error) {
	config = &DockerTestConfig{}
	if err = goaterr.ToErrors(goaterr.AppendError(nil,
		InjectEnv("DOCKER_TESTS", &config.onStr),
	)); err != nil {
		return nil, err
	}
	if config.onStr != "YES" && config.onStr != "yes" {
		return nil, fmt.Errorf("Set DOCKER_TESTS=YES environment to run docker tests")
	}
	return config, nil
}
