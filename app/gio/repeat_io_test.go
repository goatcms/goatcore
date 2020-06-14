package gio

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/filesystem"
	"github.com/goatcms/goatcore/filesystem/filespace/memfs"
)

func TestRepeatIOStory(t *testing.T) {
	t.Parallel()
	var (
		outBuf     = &bytes.Buffer{}
		in         = NewInput(strings.NewReader("some\nsecondlinetext\n"))
		out        = NewOutput(outBuf)
		cwd        filesystem.Filespace
		io         app.IO
		err        error
		firstLine  string
		secondLine string
	)
	if cwd, err = memfs.NewFilespace(); err != nil {
		t.Error(err)
		return
	}
	io = NewRepeatIO(IOParams{
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
	if !strings.Contains(outBuf.String(), "some") {
		t.Errorf("Output text should contains input line and it is '%s'", outBuf.String())
		return
	}
	if !strings.Contains(outBuf.String(), "result") {
		t.Errorf("Output text should contains writed content and it is '%s'", outBuf.String())
		return
	}
	if secondLine, err = io.In().ReadLine(); err != nil {
		t.Error(err)
		return
	}
	if secondLine != "secondlinetext" {
		t.Error("Expected second line equals to 'text'")
		return
	}
	if !strings.Contains(outBuf.String(), "secondlinetext") {
		t.Errorf("Output text should contains writed second input line and it is '%s'", outBuf.String())
		return
	}
}
