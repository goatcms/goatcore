package pipc

import (
	"strings"

	"github.com/goatcms/goatcore/varutil/goaterr"
)

func splitWaitNames(str string) (result []string, err error) {
	for _, row := range strings.Split(str, ",") {
		row = strings.Trim(row, cutset)
		if !namePattern.MatchString(row) {
			return nil, goaterr.Errorf("Incorrect name '%s'", row)
		}
		result = append(result, row)
	}
	return result, nil
}

func markBoolMapForNamespace(keys, namspace string, value bool, dest map[string]bool) (err error) {
	var (
		rows = strings.Split(keys, ",")
	)
	for _, row := range rows {
		row = strings.Trim(row, cutset)
		if strings.HasPrefix(row, "@") {
			if !namePattern.MatchString(row[1:]) {
				return goaterr.Errorf("Incorrect name '%s'", row)
			}
		} else {
			if !namePattern.MatchString(row) {
				return goaterr.Errorf("Incorrect name '%s'", row)
			}
			row = namspace + row
		}
		dest[row] = value
	}
	return nil
}
