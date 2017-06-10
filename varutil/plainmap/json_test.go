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

func TestStringPlainmapToJSON(t *testing.T) {
	t.Parallel()
	json, err := PlainStringMapToJSON(map[string]string{
		"k1.k2.k3": "value\"1\"",
		"s1.s2.s3": "value2",
		"s1.s2.s4": "value3",
	})
	if err != nil {
		t.Error(err)
		return
	}
	if json != `{"k1":{"k2":{"k3":"value\"1\""}},"s1":{"s2":{"s3":"value2","s4":"value3"}}}` {
		t.Errorf("incorrect result json: %s", json)
	}
}
