package dbconfig

import "testing"

func TestDecoderJsonData(t *testing.T) {
	decoder, err := NewDecoderFromKeyValueString(" key1=1 		 keys=s	")
	if err != nil {
		t.Error(err)
		return
	}
	// string
	if decoder.Get("key1", "0") != "1" {
		t.Error("key1 != 1")
	}
	if decoder.Get("unknowkey", "0") != "0" {
		t.Error("should return alternative string for unknow key")
	}
	// int
	if decoder.GetInt("key1", 44) != 1 {
		t.Error("key1 != 1")
	}
	if decoder.GetInt("unknowkey", 1) != 1 {
		t.Error("should return alternative int for unknow key")
	}
	if decoder.GetInt("keys", 44) != 44 {
		t.Error("keys (string value) should return alternative value")
	}
	// int64
	if decoder.GetInt64("key1", 44) != 1 {
		t.Error("key1 != 1")
	}
	if decoder.GetInt64("unknowkey", 1) != 1 {
		t.Error("should return alternative int for unknow key")
	}
	if decoder.GetInt64("keys", 44) != 44 {
		t.Error("keys (string value) should return alternative value")
	}
}
