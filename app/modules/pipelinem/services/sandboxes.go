package services

import "github.com/goatcms/goatcore/app"

// Sandbox is a security mechanism for separating running programs, usually
// in an effort to mitigate system failures or software vulnerabilities
// from spreading
type Sandbox interface {
	Run(ctx app.IOContext) (err error)
}

// SandboxBuilder is a sandbox instance generator
type SandboxBuilder interface {
	Is(name string) bool
	Build(name string) (sandbox Sandbox, err error)
}

// SandboxsManager is a tool to menage sandboxes.
type SandboxsManager interface {
	Add(sandboxsFactory SandboxBuilder)
	Get(name string) (sandbox Sandbox, err error)
}
