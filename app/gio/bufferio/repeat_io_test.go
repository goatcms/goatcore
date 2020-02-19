package bufferio

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/gio"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestRepeatIOStory(t *testing.T) {
	t.Parallel()
	var (
		buf       = &bytes.Buffer{}
		in        = gio.NewInput(strings.NewReader("some\ntext"))
		out       = gio.NewOutput(buf)
		cwd       filesystem.Filespace
		io        app.IO
		err       error
		firstLine string
	)
	if cwd, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	io = NewRepeatIO(gio.IOParams{
		In:  in,
		Out: out,
		Err: out,
		CWD: cwd,
	})
	if firstLine, err = io.In().ReadLine(); err != nil {
		t.Error(err)
		return
	}
	if firstLine != "some" {
		t.Error("Expected first line equals to 'some'")
		return
	}
	if err = io.Out().Printf("result"); err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(buf.String(), "some") {
		t.Errorf("Output text should contains input line and it is '%s'", buf.String())
		return
	}
	if !strings.Contains(buf.String(), "result") {
		t.Errorf("Output text should contains writed content and it is '%s'", buf.String())
		return
	}
}
