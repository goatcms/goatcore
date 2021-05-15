package scope

import (
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/varutil/idutil"
)

func CorrlationID(task string) (cid string) {
	if task == "" {
		task = varutil.RandString(5, varutil.AlphaNumericBytes)
	}
	return strings.Join([]string{
		idutil.CorrelationHostID(),
		strconv.FormatUint(uint64(os.Getpid()), 36),
		strconv.FormatInt(time.Now().UnixNano(), 36),
		task,
	}, "-")
}
