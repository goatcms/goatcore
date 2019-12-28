package gio

import (
	"bytes"
	"testing"

	"github.com/goatcms/goatcore/app"
)

func TestOutputBroadcast(t *testing.T) {
	t.Parallel()
	firstBuf := new(bytes.Buffer)
	firstOutput := NewAppOutput(firstBuf)
	secondBuf := new(bytes.Buffer)
	secondOutput := NewAppOutput(secondBuf)
	brodcast := NewOutputBroadcast([]app.Output{
		firstOutput,
		secondOutput,
	})
	if _, err := firstOutput.Write([]byte("0 ")); err != nil {
		t.Error(err)
		return
	}
	if err := brodcast.Printf("%d %d", 1, 2); err != nil {
		t.Error(err)
		return
	}
	if _, err := brodcast.Write([]byte(" 3")); err != nil {
		t.Error(err)
		return
	}
	result := firstBuf.String()
	if result != "0 1 2 3" {
		t.Errorf("expected '0 1 2 3' in first buffor and take %s", result)
		return
	}
	result = secondBuf.String()
	if result != "1 2 3" {
		t.Errorf("expected '1 2 3' in second buffor and take %s", result)
		return
	}
}
