package smtpmail

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/workers"
	"github.com/goatcms/goatcore/workers/jobsync"
)

func TestFormatMailContent(t *testing.T) {
	mailtime := time.Now()

	email := &goatmail.Mail{
		Date: mailtime,
		From: goatmail.Address{
			Name:    "fromname",
			Address: "from@goatcms.com",
		},
		To: []goatmail.Address{goatmail.Address{
			Name:    "toname",
			Address: "to@goatcms.com",
		}, goatmail.Address{
			Name:    "toname2",
			Address: "to2@goatcms.com",
		}},
		Subject: "Goatcore subject",
		Body: map[string]io.Reader{
			"text/plain": strings.NewReader("some plain text"),
			"text/html":  strings.NewReader("some <b>html</b>"),
		},
		Attachments: []goatmail.Attachment{
			goatmail.Attachment{
				Name:   "attachment1.txt",
				MIME:   "text/plain",
				Reader: strings.NewReader("text file content"),
			},
		},
	}

	lifecycle := jobsync.NewLifecycle(workers.DefaultTimeout, true)
	r, err := FormatMail(email, lifecycle)
	if err != nil {
		t.Error(err)
		return
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	body := buf.String()

	if len(lifecycle.Errors()) != 0 {
		t.Errorf("%v", lifecycle.Errors())
		return
	}
	fmt.Println(body)

	if !strings.Contains(body, "some plain text") {
		t.Errorf("body lost plain/text alternative")
		return
	}
	if !strings.Contains(body, "some <b>html</b>") {
		t.Errorf("body lost text/html alternative")
		return
	}
	if !strings.Contains(body, "attachment1.txt") {
		t.Errorf("body lost attachment")
		return
	}
	if !strings.Contains(body, base64.StdEncoding.EncodeToString([]byte("text file content"))) {
		t.Errorf("body lost attachment base64 content")
		return
	}

}
