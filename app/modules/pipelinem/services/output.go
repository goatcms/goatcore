package services

import "github.com/goatcms/goatcore/app"

// PipScopeOutput contains evry scope task output
type PipScopeOutput interface {
	Entrie() string
	All() map[string]string
	Writer(name string) (out app.Output)
}

// PipOutput menage pipeline outputs
type PipOutput interface {
	ScopeOutput(scope app.Scope) (output PipScopeOutput)
}
