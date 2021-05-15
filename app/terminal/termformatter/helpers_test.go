package termformatter

import (
	"testing"
)

func TestSeparateLinesAlaMa(t *testing.T) {
	line, rest := SeparateLines([]string{"ala", "ma", "kota"}, 6)
	if len(line) != 2 {
		t.Errorf("Expected [ala ma] (6 length) line and take: %v", line)
	}
	if len(rest) != 1 {
		t.Errorf("Expected [kota] as rest words and take: %v", rest)
	}
}

func TestSeparateLinesAla(t *testing.T) {
	line, rest := SeparateLines([]string{"ala", "ma", "kota"}, 5)
	if len(line) != 1 {
		t.Errorf("Expected [ala] (3 length) line and take: %v", line)
	}
	if len(rest) != 2 {
		t.Errorf("Expected [ma kota] as rest words and take: %v", rest)
	}
}

func TestSeparateLinesLongword(t *testing.T) {
	line, rest := SeparateLines([]string{"Onomatopeja", "lorem", "ipsum"}, 6)
	if len(line) != 1 {
		t.Errorf("Expected 'Onoma-' (6 length) line and take: %v", line)
		return
	}
	if line[0] != "Onoma-" {
		t.Errorf("Expected 'Onoma-' (6 length) line and take: %v", line)
	}
	if len(rest) != 3 {
		t.Errorf("Expected [-topeja lorem ipsum] as rest words and take: %v", rest)
	}
	if rest[0] != "-topeja" {
		t.Errorf("Expected '-topeja' (7 length) as first rest eleemnt: %v", rest)
	}
}

func TestSeparateLinesEntrieLine(t *testing.T) {
	line, rest := SeparateLines([]string{"a", "b", "c", "d"}, 7)
	if len(line) != 4 {
		t.Errorf("Expected [a b c d] (7 length - 4 words and 3 spaces) line and take: %v", line)
	}
	if len(rest) != 0 {
		t.Errorf("Rest is unexpected: %v", rest)
	}
}

func TestToLeft(t *testing.T) {
	result := ToLeft([]string{"a", "b", "c"}, 7)
	if result != "a b c  " {
		t.Errorf("Expected 'a b c  ' and take: %v (%d length)", result, len(result))
	}
	if len(result) != 7 {
		t.Errorf("Expected 7 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestToRight(t *testing.T) {
	result := ToRight([]string{"a", "b", "c"}, 7)
	if result != "  a b c" {
		t.Errorf("Expected '  a b c' and take: %v (%d length)", result, len(result))
	}
	if len(result) != 7 {
		t.Errorf("Expected 7 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestJustifyABC7(t *testing.T) {
	result := Justify([]string{"a", "b", "c"}, 7)
	if result != "a  b  c" {
		t.Errorf("Expected 'a  b  c' and take: %v (%d length)", result, len(result))
	}
	if len(result) != 7 {
		t.Errorf("Expected 7 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestJustifyABC8(t *testing.T) {
	result := Justify([]string{"a", "b", "c"}, 8)
	if result != "a  b   c" {
		t.Errorf("Expected 'a  b   c' and take: %v (%d length)", result, len(result))
		return
	}
	if len(result) != 8 {
		t.Errorf("Expected 8 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestJustifyABCD9(t *testing.T) {
	result := Justify([]string{"a", "b", "c", "d"}, 9)
	if result != "a b  c  d" {
		t.Errorf("Expected 'a b  c  d' and take: %v (%d length)", result, len(result))
		return
	}
	if len(result) != 9 {
		t.Errorf("Expected 9 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestJustifyABCDEF12(t *testing.T) {
	result := Justify([]string{"a", "b", "c", "d", "e", "f"}, 12)
	if result != "a b c d e  f" {
		t.Errorf("Expected 'a b c d e  f' and take: %v (%d length)", result, len(result))
		return
	}
	if len(result) != 12 {
		t.Errorf("Expected 12 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestJustifyA3(t *testing.T) {
	result := Justify([]string{"a"}, 3)
	if result != "a  " {
		t.Errorf("Expected 'a  ' and take: %v (%d length)", result, len(result))
		return
	}
	if len(result) != 3 {
		t.Errorf("Expected 3 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestJustifyAB4(t *testing.T) {
	result := Justify([]string{"a", "b"}, 4)
	if result != "a  b" {
		t.Errorf("Expected 'a  b' and take: %v (%d length)", result, len(result))
		return
	}
	if len(result) != 4 {
		t.Errorf("Expected 4 character length line and take: %v (%d length)", result, len(result))
	}
}

func TestJustifyStorySentence(t *testing.T) {
	lineWidth := 71
	result := Justify([]string{"The", "argument", "set", "path", "to", "current", "working", "directory.", "The", "CWD", "word", "is"}, lineWidth)
	expected := `The argument set  path to current  working directory. The  CWD word  is`
	if result != expected {
		t.Errorf("Expected '%s' and take: %v (%d length)", expected, result, len(result))
		return
	}
	if len(result) != lineWidth {
		t.Errorf("Expected %d character length line and take: %v (%d length)", lineWidth, result, len(result))
	}
}
