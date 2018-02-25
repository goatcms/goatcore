package totype

import "strconv"

// StringToInt convert string to int
func StringToInt(from string) (int, error) {
	val, err := strconv.ParseInt(from, DefeultNumericBase, 64)
	return int(val), err
}

// StringToInt16 convert string to int16
func StringToInt16(from string) (int16, error) {
	val, err := strconv.ParseInt(from, DefeultNumericBase, 16)
	return int16(val), err
}

// StringToInt32 convert string to int32
func StringToInt32(from string) (int32, error) {
	val, err := strconv.ParseInt(from, DefeultNumericBase, 32)
	return int32(val), err
}

// StringToInt64 convert string to int64
func StringToInt64(from string) (int64, error) {
	return strconv.ParseInt(from, DefeultNumericBase, 64)
}

// StringToUint convert string to unsigned int
func StringToUint(from string) (uint, error) {
	val, err := strconv.ParseUint(from, DefeultNumericBase, 64)
	return uint(val), err
}

// StringToUint16 convert string to unsigned int (uint16)
func StringToUint16(from string) (uint16, error) {
	val, err := strconv.ParseUint(from, DefeultNumericBase, 16)
	return uint16(val), err
}

// StringToUint32 convert string to unsigned int (uint32)
func StringToUint32(from string) (uint32, error) {
	val, err := strconv.ParseUint(from, DefeultNumericBase, 32)
	return uint32(val), err
}

// StringToUint64 convert string to unsigned int (uint64)
func StringToUint64(from string) (uint64, error) {
	return strconv.ParseUint(from, DefeultNumericBase, 64)
}

// StringToFloat32 convert string to float32
func StringToFloat32(from string) (float32, error) {
	val, err := strconv.ParseFloat(from, 32)
	return float32(val), err
}

// StringToFloat64 convert string to float64
func StringToFloat64(from string) (float64, error) {
	val, err := strconv.ParseFloat(from, 64)
	return val, err
}
