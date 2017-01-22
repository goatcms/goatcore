package totype

import "strconv"

func StringToInt(from string) (int, error) {
	val, err := strconv.ParseInt(from, DefeultNumericBase, 64)
	return int(val), err
}

func StringToInt16(from string) (int16, error) {
	val, err := strconv.ParseInt(from, DefeultNumericBase, 16)
	return int16(val), err
}

func StringToInt32(from string) (int32, error) {
	val, err := strconv.ParseInt(from, DefeultNumericBase, 32)
	return int32(val), err
}

func StringToInt64(from string) (int64, error) {
	return strconv.ParseInt(from, DefeultNumericBase, 64)
}

func StringToUint(from string) (uint, error) {
	val, err := strconv.ParseUint(from, DefeultNumericBase, 64)
	return uint(val), err
}

func StringToUint16(from string) (uint16, error) {
	val, err := strconv.ParseUint(from, DefeultNumericBase, 16)
	return uint16(val), err
}

func StringToUint32(from string) (uint32, error) {
	val, err := strconv.ParseUint(from, DefeultNumericBase, 32)
	return uint32(val), err
}

func StringToUint64(from string) (uint64, error) {
	return strconv.ParseUint(from, DefeultNumericBase, 64)
}

func StringToFloat32(from string) (float32, error) {
	val, err := strconv.ParseFloat(from, 32)
	return float32(val), err
}

func StringToFloat64(from string) (float64, error) {
	val, err := strconv.ParseFloat(from, 64)
	return val, err
}
