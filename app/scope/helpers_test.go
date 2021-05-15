package scope

import (
	"testing"

	"github.com/goatcms/goatcore/app"
	"github.com/goatcms/goatcore/app/scope/datascope"
)

func TestGetString(t *testing.T) {
	var (
		err   error
		value string
		scp   app.DataScope
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": "value",
	})
	if value, err = GetString(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != "value" {
		t.Errorf("expected value equals to 'value' and take %s", value)
	}
}

func TestGetBool(t *testing.T) {
	var (
		scp   app.DataScope
		value bool
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": true,
	})
	if value, err = GetBool(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != true {
		t.Errorf("expected value equals to true and take %v", value)
	}
}

func TestGetInt(t *testing.T) {
	var (
		scp   app.DataScope
		value int
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": 1991,
	})
	if value, err = GetInt(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 1991 {
		t.Errorf("expected value equals to 1991 and take %v", value)
	}
}

func TestGetInt64(t *testing.T) {
	var (
		scp   app.DataScope
		value int64
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": int64(1991),
	})
	if value, err = GetInt64(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 1991 {
		t.Errorf("expected value equals to 1991 and take %v", value)
	}
}

func TestGetUint(t *testing.T) {
	var (
		scp   app.DataScope
		value uint
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": uint(1991),
	})
	if value, err = GetUint(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 1991 {
		t.Errorf("expected value equals to 1991 and take %v", value)
	}
}

func TestGetUint64(t *testing.T) {
	var (
		scp   app.DataScope
		value uint64
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": uint64(1991),
	})
	if value, err = GetUint64(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 1991 {
		t.Errorf("expected value equals to 1991 and take %v", value)
	}
}

func TestGetFloat32(t *testing.T) {
	var (
		scp   app.DataScope
		value float32
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": float32(6.9),
	})
	if value, err = GetFloat32(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 6.9 {
		t.Errorf("expected value equals to 6.9 and take %v", value)
	}
}

func TestGetFloat64(t *testing.T) {
	var (
		scp   app.DataScope
		value float64
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": float64(6.9),
	})
	if value, err = GetFloat64(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 6.9 {
		t.Errorf("expected value equals to 6.9 and take %v", value)
	}
}

func TestGetComplex64(t *testing.T) {
	var (
		scp   app.DataScope
		value complex64
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": complex64(6.9),
	})
	if value, err = GetComplex64(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 6.9 {
		t.Errorf("expected value equals to 6.9 and take %v", value)
	}
}

func TestGetComplex128(t *testing.T) {
	var (
		scp   app.DataScope
		value complex128
		err   error
	)
	t.Parallel()
	scp = datascope.New(map[interface{}]interface{}{
		"key": complex128(6.9),
	})
	if value, err = GetComplex128(scp, "key"); err != nil {
		t.Error(err)
		return
	}
	if value != 6.9 {
		t.Errorf("expected value equals to 6.9 and take %v", value)
	}
}
