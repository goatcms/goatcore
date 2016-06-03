package testbase

import (
	"github.com/goatcms/goat-core/varutil"
)

func ByteArrayEq(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func ComparePublicEquals(obj1 interface{}, obj2 interface{}) (bool, error) {
	json1, err := varutil.ObjectToJson(obj1)
	if err != nil {
		return false, err
	}
	json2, err := varutil.ObjectToJson(obj2)
	if err != nil {
		return false, err
	}
  return json1 == json2, nil
}
