package testbase

import (
	"time"

	"github.com/goatcms/goatcore/varutil/goaterr"

	"github.com/goatcms/goatcore/goatnet"
)

// ReadURLLoop read url in loop. It is used for integration tests.
func ReadURLLoop(url string, count int, dur time.Duration) (respBody string, err error) {
	var (
		respBytes []byte
		errs      []error
	)
	for ; count > 0; count-- {
		time.Sleep(dur)
		if respBytes, err = goatnet.ReadURL(url); err == nil {
			return string(respBytes), nil
		}
		errs = append(errs, err)
	}
	return "", goaterr.ToError(errs)
}
