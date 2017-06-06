package plainmap

import "testing"

const (
	jsonplainmapdata = `{"stringkey":"stringvalue", "object":{"stringkey2":"stringvalue2"}, "number":11}`
)

func TestJSONToPlainStringMap(t *testing.T) {
	t.Parallel()
	outmap, err := JSONToPlainStringMap([]byte(jsonplainmapdata))
	if err != nil {
		t.Error(err)
		return
	}

	if outmap["stringkey"] != "stringvalue" {
		t.Error("stringkey != stringvalue")
	}
	if outmap["object.stringkey2"] != "stringvalue2" {
		t.Error("object.stringkey2 != stringvalue2")
	}
	if outmap["number"] != "11" {
		t.Error("number != 11")
	}
}
