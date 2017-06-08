package gio

import (
	"io"
	"strings"
	"testing"
)

const (
	manySpaces    = "                                    "
	manyTabs      = "\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"
	testSentence  = manySpaces + manyTabs + "line1.1  line1.2 \n line2.1 " + manySpaces + manyTabs + " line2.2 "
	testSentence2 = "alamakota"
)

func TestReadWord(t *testing.T) {
	t.Parallel()
	var value string
	var err error
	rd := strings.NewReader(testSentence)
	input := NewInput(rd)

	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line1.1" {
		t.Errorf("expected line1.1 and take %s", value)
		return
	}
	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line1.2" {
		t.Errorf("expected line1.2 and take %s", value)
		return
	}
	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line2.1" {
		t.Errorf("expected lin2e.1 and take %s", value)
		return
	}
	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line2.2" {
		t.Errorf("expected line2.2 and take %s", value)
		return
	}
	value, err = input.ReadWord()
	if err != io.EOF {
		t.Errorf("expected io.EOF error and take %v", err)
		return
	}
	if value != "" {
		t.Errorf("expected empty string and take %s", value)
		return
	}
}

func TestReadWord2(t *testing.T) {
	t.Parallel()
	var value string
	var err error
	rd := strings.NewReader(testSentence2)
	input := NewInput(rd)
	value, err = input.ReadWord()
	if err != io.EOF {
		t.Errorf("expected io.EOF error and take %v", err)
		return
	}
	if value != "alamakota" {
		t.Errorf("expected alamakota and take %s", value)
		return
	}
}

func TestReadWordSized(t *testing.T) {
	t.Parallel()
	var value string
	var err error
	rd := strings.NewReader(testSentence)
	input := NewInputSize(rd, 11)

	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line1.1" {
		t.Errorf("expected line1.1 and take %s", value)
		return
	}
	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line1.2" {
		t.Errorf("expected line1.2 and take %s", value)
		return
	}
	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line2.1" {
		t.Errorf("expected lin2e.1 and take %s", value)
		return
	}
	if value, err = input.ReadWord(); err != nil {
		t.Error(err)
		return
	}
	if value != "line2.2" {
		t.Errorf("expected line2.2 and take %s", value)
		return
	}
}

func TestReadLine(t *testing.T) {
	t.Parallel()
	var value string
	var err error
	rd := strings.NewReader(testSentence)
	input := NewInput(rd)

	if value, err = input.ReadLine(); err != nil {
		t.Error(err)
		return
	}
	if value != "line1.1  line1.2 " {
		t.Errorf("expected first line without a apaces at begin and take '%s'", value)
		return
	}
	value, err = input.ReadLine()
	if err != io.EOF {
		t.Errorf("expected io.EOF error and take %v", err)
		return
	}
	if value != "line2.1 "+manySpaces+manyTabs+" line2.2 " {
		t.Errorf("expected second line without a apaces at begin and take '%s'", value)
		return
	}
	if _, err = input.ReadLine(); err != io.EOF {
		t.Errorf("expected io.EOF error")
		return
	}
}

func TestReadLine2(t *testing.T) {
	t.Parallel()
	var value string
	var err error
	rd := strings.NewReader(testSentence2)
	input := NewInput(rd)
	value, err = input.ReadLine()
	if err != io.EOF {
		t.Errorf("expected io.EOF error and take %v", err)
		return
	}
	if value != "alamakota" {
		t.Errorf("expected alamakota and take %s", value)
		return
	}
}
