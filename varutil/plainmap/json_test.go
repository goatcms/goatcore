package plainmap

import (
	"strings"
	"testing"
)

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

func TestStringPlainmapToFormattedJSON(t *testing.T) {
	t.Parallel()
	result, err := PlainStringMapToFormattedJSON(map[string]string{
		"k1.k2.k3": "value\"1\"",
		"s1.s2.s3": "value2",
		"s1.s2.s4": "value3",
	})
	if err != nil {
		t.Error(err)
		return
	}
	s2 := "  "
	s4 := s2 + s2
	s6 := s4 + s2
	expected := "{\n" + s2 + "\"k1\": {\n" + s4 + "\"k2\": {\n" + s6 + "\"k3\": \"value\\\"1\\\"\"\n" + s4 + "}\n" + s2 + "},\n" + s2 + "\"s1\": {\n" + s4 + "\"s2\": {\n" + s6 + "\"s3\": \"value2\",\n" + s6 + "\"s4\": \"value3\"\n" + s4 + "}\n" + s2 + "}\n}"
	expectedLines := strings.Split(expected, "\n")
	resultLines := strings.Split(result, "\n")
	if len(expectedLines) != len(resultLines) {
		t.Errorf("expected the same line numbers (result has %v lines, expected is %v", len(resultLines), len(expectedLines))
	}
	for i, _ := range resultLines {
		if expectedLines[i] != resultLines[i] {
			t.Errorf("line %v is diffrent expected '%v' and take '%v'", i, expectedLines[i], resultLines[i])
		}
	}
}
