package goatapp

import (
	"bytes"
	"io"
	"strings"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

// MockupApp is default mockup applicationo (used by test code)
type MockupApp struct {
	app.App
	outBuf bytes.Buffer
	errBuf bytes.Buffer
}

// NewMockupApp create new mockup application
func NewMockupApp(params Params) (mockup *MockupApp, err error) {
	mockup = &MockupApp{}
	if params.Name == "" {
		params.Name = "TestApp"
	}
	if params.Arguments == nil {
		params.Arguments = []string{}
	}
	if params.IO.Err == nil {
		params.IO.Err = gio.NewAppOutput(&mockup.outBuf)
	} else {
		params.IO.Err = gio.NewAppOutput(io.MultiWriter(
			&mockup.outBuf,
			params.IO.Err,
		))
	}
	if params.IO.In == nil {
		params.IO.In = gio.NewAppInput(strings.NewReader(""))
	}
	if params.IO.Out == nil {
		params.IO.Out = gio.NewAppOutput(&mockup.outBuf)
	} else {
		params.IO.Out = gio.NewAppOutput(io.MultiWriter(
			&mockup.outBuf,
			params.IO.Out,
		))
	}
	if params.Filespaces.Root == nil {
		if params.Filespaces.Root, err = memfs.NewFilespace(); err != nil {
			return
		}
	}
	if params.Filespaces.CWD == nil {
		if err = params.Filespaces.Root.MkdirAll("cwd", 0766); err != nil {
			return
		}
		if params.Filespaces.CWD, err = params.Filespaces.Root.Filespace("cwd"); err != nil {
			return
		}
	}
	if params.Filespaces.Home == nil {
		if err = params.Filespaces.Root.MkdirAll("home/username", 0766); err != nil {
			return nil, err
		}
		if params.Filespaces.Home, err = params.Filespaces.Root.Filespace("home/username"); err != nil {
			return nil, err
		}
	}
	if params.Filespaces.Tmp == nil {
		if err = params.Filespaces.Root.MkdirAll("tmp", 0766); err != nil {
			return
		}
		if params.Filespaces.Tmp, err = params.Filespaces.Root.Filespace("tmp"); err != nil {
			return
		}
	}
	if mockup.App, err = NewGoatApp(params); err != nil {
		return
	}
	return mockup, nil
}

// OutputBuffer return output buffer
func (mapp *MockupApp) OutputBuffer() *bytes.Buffer {
	return &mapp.outBuf
}

// ErrorBuffer return error output buffer
func (mapp *MockupApp) ErrorBuffer() *bytes.Buffer {
	return &mapp.errBuf
}
