package gio

import (
	"bytes"
	"testing"
)

func TestOutput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	output := NewAppOutput(buf)
	output.Printf("%d %d %d", 1, 2, 3)
	result := buf.String()
	if result != "1 2 3" {
		t.Errorf("expected '1 2 3' and take %s", result)
		return
	}
}
