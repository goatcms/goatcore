package abstracttype

import (
	"io/ioutil"

	"github.com/goatcms/goat-core/types"
)

// StringFromMultipart read all data from multipart file
func StringFromMultipart(fh types.FileHeader) (string, error) {
	f, err := fh.Open()
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
