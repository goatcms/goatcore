package bufferio

import (
	"bytes"
	"io"
	"testing"

	"github.com/goatcms/goatcore/app"
)

func TestMultiWriteToBrodcast(t *testing.T) {
	t.Parallel()
	firstBuf := new(bytes.Buffer)
	secondBuf := new(bytes.Buffer)
	brodcast := NewBroadcast(nil, []io.Writer{
		firstBuf,
		secondBuf,
	})
	if _, err := firstBuf.Write([]byte("0 ")); err != nil {
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

func TestAddNewOutputToBrodcast(t *testing.T) {
	var (
		err    error
		result string
	)
	t.Parallel()
	buf := new(bytes.Buffer)
	brodcast := NewBroadcast(nil, []io.Writer{})
	if _, err := brodcast.Write([]byte("1 2 3 4")); err != nil {
		t.Error(err)
		return
	}
	if err = brodcast.Add(buf); err != nil {
		t.Error(err)
		return
	}
	result = buf.String()
	if result != "1 2 3 4" {
		t.Errorf("expected '1 2 3 4 in second buffor and take %s", result)
		return
	}
}

func TestBrodcastBuffer(t *testing.T) {
	var result string
	t.Parallel()
	brodcast := NewBroadcast(nil, []io.Writer{})
	if _, err := brodcast.Write([]byte("1 2 3 4")); err != nil {
		t.Error(err)
		return
	}
	result = brodcast.String()
	if result != "1 2 3 4" {
		t.Errorf("expected '1 2 3 4 in second buffor and take %s", result)
		return
	}
}

func testBrodcastImplementsAppOutput() app.Output {
	return NewBroadcast(nil, []io.Writer{})
}
