package varutil

import "testing"

func TestQuoteReturnNull(t *testing.T) {
	t.Parallel()
	result := Quote(nil)
	if result != "null" {
		t.Errorf("Quote should return 'null' string for nil values. Result is %v and expected %v", result, "null")
		return
	}
}

func TestQuoteReturnQuotedValue(t *testing.T) {
	t.Parallel()
	v := "1234"
	result := Quote(&v)
	if result != `"1234"` {
		t.Errorf("Quote should return \"1234\" and the result is %v", result)
		return
	}
}

func TestFormatIntReturnNull(t *testing.T) {
	t.Parallel()
	result := FormatInt(nil, 10)
	if result != "null" {
		t.Errorf("Quote should return 'null' string for nil values. Result is %v and expected %v", result, "null")
		return
	}
}

func TestFormatIntReturnValue(t *testing.T) {
	t.Parallel()
	v := int64(12345)
	result := FormatInt(&v, 10)
	if result != `12345` {
		t.Errorf("Quote should return \"12345\" and the result is %v", result)
		return
	}
}

func TestQuoteArray(t *testing.T) {
	t.Parallel()
	v := []string{"1", "2", "3", "4"}
	result := QuoteArray(v, ", ")
	if result != `"1", "2", "3", "4"` {
		t.Errorf("Quote should return \"1\", \"2\", \"3\", \"4\" and the result is %v", result)
		return
	}
}

func TestFormatIntArray(t *testing.T) {
	t.Parallel()
	v := []int64{1, 2, 3, 4}
	result := FormatIntArray(v, 10, ", ")
	if result != `1, 2, 3, 4` {
		t.Errorf("Quote should return 1, 2, 3, 4 and the result is %v", result)
		return
	}
}
