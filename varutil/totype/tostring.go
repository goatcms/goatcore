package totype

import (
	"strconv"

	"github.com/goatcms/goatcore/varutil"
)

// BoolToString convert boolean to string
func BoolToString(from bool) (string, error) {
	return strconv.FormatBool(from), nil
}

// IntToString convert int to string
func IntToString(from int) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

// Int16ToString convert int16 to string
func Int16ToString(from int16) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

// Int32ToString convert int32 to string
func Int32ToString(from int32) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

// Int64ToString convert int64 to string
func Int64ToString(from int64) (string, error) {
	return strconv.FormatInt(int64(from), DefeultNumericBase), nil
}

// UintToString convert unsigned int to string
func UintToString(from uint) (string, error) {
	return strconv.FormatUint(uint64(from), DefeultNumericBase), nil
}

// Uint16ToString convert unsigned int (uint16) to string
func Uint16ToString(from uint16) (string, error) {
	return strconv.FormatUint(uint64(from), DefeultNumericBase), nil
}

// Uint32ToString convert unsigned int (uint32) to string
func Uint32ToString(from uint32) (string, error) {
	return strconv.FormatUint(uint64(from), DefeultNumericBase), nil
}

// Uint64ToString convert unsigned int (uint64) to string
func Uint64ToString(from uint64) (string, error) {
	return strconv.FormatUint(from, DefeultNumericBase), nil
}

// Float32ToString convert float32 to string
func Float32ToString(from float32) (string, error) {
	return strconv.FormatFloat(float64(from), 'E', -1, 32), nil
}

// Float64ToString convert float64 to string
func Float64ToString(from float64) (string, error) {
	return strconv.FormatFloat(from, 'E', -1, 32), nil
}

// ToString convert object to JSON string
func ToString(from interface{}) (string, error) {
	return varutil.ObjectToJSON(from)
}
