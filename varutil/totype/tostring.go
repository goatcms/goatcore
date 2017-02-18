package totype

import (
	"strconv"

	"github.com/goatcms/goatcore/varutil"
)

func BoolToString(from bool) (string, error) {
	return strconv.FormatBool(from), nil
}

func IntToString(from int) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

func Int16ToString(from int16) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

func Int32ToString(from int32) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

func Int64ToString(from int64) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

func UintToString(from uint) (string, error) {
	return strconv.FormatUint(uint64(from), DefeultNumericBase), nil
}

func Uint16ToString(from uint16) (string, error) {
	return strconv.FormatUint(uint64(from), DefeultNumericBase), nil
}

func Uint32ToString(from uint32) (string, error) {
	return strconv.FormatUint(uint64(from), DefeultNumericBase), nil
}

func Uint64ToString(from uint64) (string, error) {
	return strconv.FormatUint(from, DefeultNumericBase), nil
}

func Float32ToString(from float32) (string, error) {
	return strconv.FormatFloat(float64(from), 'E', -1, 32), nil
}

func Float64ToString(from float64) (string, error) {
	return strconv.FormatFloat(from, 'E', -1, 32), nil
}

func ToString(from interface{}) (string, error) {
	return varutil.ObjectToJSON(from)
}
