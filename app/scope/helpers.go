package scope

import (
	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/varutil/goaterr"
)

// GetString get string value from data scope
func GetString(scp app.DataScope, name string) (value string, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(string); !ok {
		return "", goaterr.Errorf("%v %T is not a string", ins, ins)
	}
	return value, nil
}

// GetBool get bool value from data scope
func GetBool(scp app.DataScope, name string) (value bool, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(bool); !ok {
		return false, goaterr.Errorf("%v %T is not a bool", ins, ins)
	}
	return value, nil
}

// GetInt get int value from data scope
func GetInt(scp app.DataScope, name string) (value int, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(int); !ok {
		return value, goaterr.Errorf("%v %T is not a int", ins, ins)
	}
	return value, nil
}

// GetInt64 get int value from data scope
func GetInt64(scp app.DataScope, name string) (value int64, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(int64); !ok {
		return value, goaterr.Errorf("%v %T is not a int64", ins, ins)
	}
	return value, nil
}

// GetUint get uint value from data scope
func GetUint(scp app.DataScope, name string) (value uint, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(uint); !ok {
		return value, goaterr.Errorf("%v %T is not a uint", ins, ins)
	}
	return value, nil
}

// GetUint64 get uint64 value from data scope
func GetUint64(scp app.DataScope, name string) (value uint64, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(uint64); !ok {
		return value, goaterr.Errorf("%v %T is not a uint64", ins, ins)
	}
	return value, nil
}

// GetFloat32 get float32 value from data scope
func GetFloat32(scp app.DataScope, name string) (value float32, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(float32); !ok {
		return value, goaterr.Errorf("%v %T is not a float32", ins, ins)
	}
	return value, nil
}

// GetFloat64 get float64 value from data scope
func GetFloat64(scp app.DataScope, name string) (value float64, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(float64); !ok {
		return value, goaterr.Errorf("%v %T is not a float64", ins, ins)
	}
	return value, nil
}

// GetComplex64 get complex64 value from data scope
func GetComplex64(scp app.DataScope, name string) (value complex64, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(complex64); !ok {
		return value, goaterr.Errorf("%v %T is not a complex64", ins, ins)
	}
	return value, nil
}

// GetComplex128 get complex128 value from data scope
func GetComplex128(scp app.DataScope, name string) (value complex128, err error) {
	var (
		ins = scp.Value(name)
		ok  bool
	)
	if value, ok = ins.(complex128); !ok {
		return value, goaterr.Errorf("%v %T is not a complex64", ins, ins)
	}
	return value, nil
}
