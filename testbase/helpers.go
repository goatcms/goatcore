package testbase

import (
	"os"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

func InjectEnv(name string, dest *string) (err error) {
	if *dest = os.Getenv(name); *dest == "" {
		return goaterr.Errorf("%v operating environment is required", name)
	}
	return nil
}
