package gio

import "testing"

func TestIsWhiteChar(t *testing.T) {
	t.Parallel()
	if !isWhiteChar(' ') {
		t.Errorf("space is white char")
	}
	if !isWhiteChar('\t') {
		t.Errorf("tab is white char")
	}
	if !isWhiteChar('\r') {
		t.Errorf("\\r is white char")
	}
	if !isWhiteChar('\n') {
		t.Errorf("\\n is white char")
	}
}

func TestIsNewLine(t *testing.T) {
	t.Parallel()
	if !isNewLine('\r') {
		t.Errorf("\\r is new line ascii")
	}
	if !isNewLine('\n') {
		t.Errorf("\\n is new line ascii")
	}
}
