package termformatter

import (
	"bytes"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/app/gio"
)

func TestBlockFormatterStory(t *testing.T) {
	var (
		buf = bytes.NewBuffer([]byte{})
		out = gio.NewAppOutput(buf)
	)
	formatter := NewBlockFormatter(out, 15,
		NewBlockDef(5, ToRight, ToRight),
		NewBlockDef(10, Justify, ToLeft),
	)
	formatter.PrintBlocks(
		"a b c",
		"12 34 56 78",
	)
	resultLines := strings.Split(buf.String(), "\n")
	if len(resultLines) != 3 {
		t.Errorf("Expected three result lines and teke %d lines:\n`%s`", len(resultLines), buf.String())
	}
	expected := `a b c12  34  56`
	if resultLines[0] != expected {
		t.Errorf("Expected first line as `a b c12  34  56` and take:\n`%s`", buf.String())
	}
	if len(resultLines[0]) != 15 {
		t.Errorf("Expected first line 15 chars length")
	}
	expected = `     78        `
	if resultLines[1] != expected {
		t.Errorf("Expected second line as `     78        ` and take:\n`%s`", buf.String())
	}
	if len(resultLines[1]) != 15 {
		t.Errorf("Expected second line 15 chars length")
	}
	if resultLines[2] != "" {
		t.Errorf("Last line must be empty")
	}
}
