package testbase

import (
	"fmt"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

// PodmanTestConfig is podman test config object
type PodmanTestConfig struct {
	onStr string
}

// LoadPodmanTestConfig return docker config from envs
func LoadPodmanTestConfig() (config *PodmanTestConfig, err error) {
	config = &PodmanTestConfig{}
	if err = goaterr.ToError(goaterr.AppendError(nil,
		InjectEnv("GOATCORE_TEST_PODMAN", &config.onStr),
	)); err != nil {
		return nil, err
	}
	if config.onStr != "YES" && config.onStr != "yes" {
		return nil, fmt.Errorf("export GOATCORE_TEST_DOCKER=YES environment to run docker tests")
	}
	return config, nil
}
