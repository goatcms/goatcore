package ocservices

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/filesystem"
)

const (
	// DockerEngine is a identifier for docker container system ( https://www.docker.com/ )
	DockerEngine = "docker"
	// PodmanEngine is a identifier for podman container system ( https://podman.io/ )
	PodmanEngine = "podman"
	// DefaultEngine is a identifier if default container system
	DefaultEngine = PodmanEngine
)

// FSVolume represent single volume mapping
type FSVolume struct {
	Filespace filesystem.Filespace
	Path      string
}

// Container describe execution task
type Container struct {
	IO         app.IO
	Image      string
	WorkDir    string
	Entrypoint string
	Envs       commservices.Environments
	Engine     string
	Provilages bool
	FSVolumes  map[string]FSVolume
}

// Engine provide container engine wrapper
type Engine interface {
	// Run new temporary container
	Run(container Container) error
}

// Manager provide container system
type Manager interface {
	// AddEngine register new engine
	AddEngine(id string, engine Engine) (err error)
	SetDefaultEngine(engine Engine)
	Engine
}
