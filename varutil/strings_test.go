package varutil

import (
	"testing"
)

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
