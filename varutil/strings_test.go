package varutil

import (
	"testing"
)

func TestSplitArguments(t *testing.T) {
	var (
		err  error
		args []string
	)
	t.Parallel()
	if args, err = SplitArguments(`   v1 path=my/path numbers="12 \\"12\\"" "--some=true true" quotationMarks=\"   `); err != nil {
		t.Error(err)
		return
	}
	checkSplitArguments(t, args, 0, "v1")
	checkSplitArguments(t, args, 1, "path=my/path")
	checkSplitArguments(t, args, 2, "numbers=12 \"12\"")
	checkSplitArguments(t, args, 3, "--some=true true")
	checkSplitArguments(t, args, 4, "quotationMarks=\"")
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

func TestUnescapeString(t *testing.T) {
	expect := "some user's caffe named \"earl coffe\"\nTake here from big city."
	if take := UnescapeString("some user\\'s caffe named \\\"earl coffe\\\"\\nTake here from big city."); take != expect {
		t.Errorf("take '%s' and expect '%s'", take, expect)
	}
}

func TestUnescapeStringSpecialChars(t *testing.T) {
	expect := "\n\t\"'"
	if take := UnescapeString("\\n\\t\\\"'"); take != expect {
		t.Errorf("take '%s' and expect '%s'", take, expect)
	}
}
