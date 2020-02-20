package smtpmailexample

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/goatcms/goatcore/goatmail/smtpmail"
)

func EncodeString(s string) (result string, err error) {
	var buff []byte
	encoder := smtpmail.NewBase64Encoder(strings.NewReader(s))
	if buff, err = ioutil.ReadAll(encoder); err != nil {
		return "", err
	}
	return string(buff), nil
}

func TestEncodeString(t *testing.T) {
	var (
		err    error
		result string
	)
	t.Parallel()
	if result, err = EncodeString("Hello World!"); err != nil {
		t.Error(err)
		return
	}
	if result != "SGVsbG8gV29ybGQh" {
		t.Errorf("take '%v' and expected 'SGVsbG8gV29ybGQh'", result)
	}
}
