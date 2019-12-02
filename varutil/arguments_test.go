package varutil

import (
	"testing"
)

func TestStringWithWhiteSpacesIsEmptyArguments(t *testing.T) {
	var (
		eof  bool
		err  error
		args []string
	)
	t.Parallel()
	if args, eof, err = SplitArguments("\t \n \t\t"); err != nil {
		t.Error(err)
		return
	}
	if len(args) != 0 {
		t.Errorf("Expected empty arguments")
		return
	}
	if eof != false {
		t.Errorf("It contains two separated lines. Read and parse firs one. Expected eof equals to false")
		return
	}
}

func TestEmptyStrinsIsEmptyArguments(t *testing.T) {
	var (
		eof  bool
		err  error
		args []string
	)
	t.Parallel()
	if args, eof, err = SplitArguments(""); err != nil {
		t.Error(err)
		return
	}
	if len(args) != 0 {
		t.Errorf("Expected empty arguments")
		return
	}
	if eof != true {
		t.Errorf("Expected eof")
		return
	}
}

func TestSplitArguments(t *testing.T) {
	var (
		eof  bool
		err  error
		args []string
	)
	t.Parallel()
	if args, eof, err = SplitArguments(`   v1 path=my/path numbers="12 \"12\"" "--some=true true" backslash=\\ quotationMarks=\"   `); err != nil {
		t.Error(err)
		return
	}
	if eof != true {
		t.Errorf("Expected eof")
		return
	}
	checkSplitArguments(t, args, 0, "v1")
	checkSplitArguments(t, args, 1, "path=my/path")
	checkSplitArguments(t, args, 2, "numbers=12 \"12\"")
	checkSplitArguments(t, args, 3, "--some=true true")
	checkSplitArguments(t, args, 4, "backslash=\\")
	checkSplitArguments(t, args, 5, "quotationMarks=\"")
}

func TestSkipNewLineArguments(t *testing.T) {
	var (
		eof  bool
		err  error
		args []string
	)
	t.Parallel()
	if args, eof, err = SplitArguments(`   command arg1 \
		arg2
		skipped line   `); err != nil {
		t.Error(err)
		return
	}
	if eof != false {
		t.Errorf("It contains 'skipped line' after comand. Expected eof equals to false")
		return
	}
	if len(args) != 3 {
		t.Errorf("Expected two arguments and take %v: %v", len(args), args)
		return
	}
}

func TestEscepeNewLine(t *testing.T) {
	var (
		eof  bool
		err  error
		args []string
	)
	t.Parallel()
	if args, eof, err = SplitArguments(`   v1\
		path=my/path \
		numbers="12 \\"12\\"" \
		 "--some=true true"\
     quotationMarks=\"   `); err != nil {
		t.Error(err)
		return
	}
	if eof != true {
		t.Errorf("Expected eof")
		return
	}
	checkSplitArguments(t, args, 0, "v1")
	checkSplitArguments(t, args, 1, "path=my/path")
	checkSplitArguments(t, args, 2, "numbers=12 \"12\"")
	checkSplitArguments(t, args, 3, "--some=true true")
	checkSplitArguments(t, args, 4, "quotationMarks=\"")
}

func TestSplitMultilineArgument(t *testing.T) {
	var (
		eof  bool
		err  error
		args []string
	)
	t.Parallel()
	if args, eof, err = SplitArguments(`   v1 path=<<EOF
		Some
		Multiline
		argument
EOF`); err != nil {
		t.Error(err)
		return
	}
	if eof != true {
		t.Errorf("Expected eof")
		return
	}
	checkSplitArguments(t, args, 0, "v1")
	checkSplitArguments(t, args, 1, `path=Some
		Multiline
		argument`)
}

func checkSplitArguments(t *testing.T, args []string, index int, expected string) {
	var have string
	if index >= len(args) {
		t.Errorf("have %d arguments. Argument %d doesn't exist", len(args), index)
		return
	}
	have = args[index]
	if have != expected {
		t.Errorf(`element by %d index must be equal to "%v" and take "%v"`, index, expected, have)
	}
}
