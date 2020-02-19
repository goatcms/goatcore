package gio

import (
	"bytes"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func newEmptyIO() (ctx app.IO) {
	var (
		cwd filesystem.Filespace
		err error
	)
	in := NewInput(new(bytes.Buffer))
	out := NewOutput(new(bytes.Buffer))
	if cwd, err = memfs.NewFilespace(); err != nil {
		panic(err)
	}
	return NewIO(IOParams{
		In:  in,
		Out: out,
		Err: out,
		CWD: cwd,
	})
}
