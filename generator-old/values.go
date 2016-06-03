package generator

import (
	"github.com/goatcms/goat-core/varutil"
)

type Values map[string]string

func (v *Values) Read(path string) error {
	return varutil.ReadJson(path, v)
}

func (v *Values) Write(path string) error {
	return varutil.WriteJson(path, v)
}
