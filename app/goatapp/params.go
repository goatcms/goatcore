package goatapp

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/dependency"
	"github.com/goatcms/goatcore/filesystem"
)

type Filespaces struct {
	CWD  filesystem.Filespace
	Home filesystem.Filespace
	Root filesystem.Filespace
	Tmp  filesystem.Filespace
}

type filespacesProvider struct {
	fs Filespaces
}

func (provider filespacesProvider) CWD() filesystem.Filespace {
	return provider.fs.CWD
}

func (provider filespacesProvider) Home() filesystem.Filespace {
	return provider.fs.Home
}

func (provider filespacesProvider) Root() filesystem.Filespace {
	return provider.fs.Root
}

func (provider filespacesProvider) Tmp() filesystem.Filespace {
	return provider.fs.Tmp
}

type Scopes struct {
	App       app.Scope
	args      app.DataScope
	config    app.DataScope
	filespace app.DataScope
}

type scopesProvider struct {
	scopes Scopes
}

func (provider scopesProvider) App() app.Scope {
	return provider.scopes.App
}

func (provider scopesProvider) Arguments() app.DataScope {
	return provider.scopes.args
}

func (provider scopesProvider) Config() app.DataScope {
	return provider.scopes.config
}

func (provider scopesProvider) Filespace() app.DataScope {
	return provider.scopes.filespace
}

type IO struct {
	In  app.Input
	Out app.Output
	Err app.Output
}

// Params define new app
type Params struct {
	Arguments      []string
	DP             dependency.Provider
	Env            string
	Filespaces     Filespaces
	HealthCheckers app.AppHealthCheckers
	IO             IO
	Name           string
	Scopes         Scopes
	Terminal       app.TerminalManager
	Version        app.Version
}
