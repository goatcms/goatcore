package pipservices

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/modules/commonm/commservices"
	"github.com/goatcms/goatcore/filesystem"
)

// PipContext contains pip I/O data
type PipContext struct {
	In    app.Input
	Out   app.Output
	Err   app.Output
	CWD   filesystem.Filespace
	Scope app.Scope
}

// Pip describe single commands pipeline
type Pip struct {
	Context PipContext
	Name    string
	Sandbox string
	Lock    commservices.LockMap
	Wait    []string
}
