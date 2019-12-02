package workers

import (
	"fmt"
	"time"

	"github.com/goatcms/goatcore/varutil"
)

// DateTimeStr convert time.Time to formated string
func DateTimeStr(t time.Time) string {
	return fmt.Sprintf("%v-%d-%v-%v:%v:%v:%v", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}

// TimeStr convert time.Time to formated string
func TimeStr(t time.Time) string {
	return fmt.Sprintf("%v:%v:%v:%v", t.Hour(), t.Minute(), t.Second(), t.Nanosecond())
}

// NewResponseID create a new ID for response
func NewResponseID(taskName string) string {
	var (
		now  = time.Now()
		salt string
	)
	salt = varutil.RandString(4, varutil.AlphaNumericBytes)
	return DateTimeStr(now) + "_" + salt + "_" + taskName
}
