package definition

import (
	"testing"
)

func TestParamsDecoding(t *testing.T) {
	var p Params = Params{"p1", "p2", "p3", "name1:value1", "name2:value2"}
	args := p.Args()
	if args[0] != "p1" {
		t.Error("args get first value error")
	}
	if args[2] != "p3" {
		t.Error("args get first value error")
	}
	if args[3] != "name1:value1" {
		t.Error("args get first value error")
	}
	if p.Key("name1") != "value1" {
		t.Error("get by name1 fail. Get " + p.Key("name1") + ", expect value1")
	}
	if p.Key("name2") != "value2" {
		t.Error("get by name2 fail. Get " + p.Key("name2") + ", expect value1")
	}
	if p.Key("name3") != "" {
		t.Error("get ba name1 fail")
	}
}
