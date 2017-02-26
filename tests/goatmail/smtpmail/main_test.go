package smtpmail_test

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/goatmail/smtpmail"
	"github.com/goatcms/goatcore/workers"
	"github.com/goatcms/goatcore/workers/jobsync"
)

func TestSendEmail(t *testing.T) {
	config, err := LoadTestConfig()
	if err != nil {
		t.Error(err)
		return
	}

	sender := smtpmail.NewMailSender(config.SenderConfig)

	mailtime := time.Now()
	mail := &goatmail.Mail{
		Date: mailtime,
		From: goatmail.Address{
			Name:    config.FromAddress,
			Address: config.FromAddress,
		},
		To: []goatmail.Address{goatmail.Address{
			Name:    config.ToAddress,
			Address: config.ToAddress,
		}},
		Subject: "Goatcore subject",
		Body: map[string]io.Reader{
			"text/plain": strings.NewReader("some content"),
			"text/html":  strings.NewReader("some <b>content</b>"),
		},
		Attachments: []goatmail.Attachment{
			goatmail.Attachment{
				Name:   "attachment1.txt",
				MIME:   "text/plain",
				Reader: strings.NewReader("text file content"),
			},
		},
	}

	lc := jobsync.NewLifecycle(workers.DefaultTestTimeout, true)
	if err := sender.Send(mail, lc); err != nil {
		t.Error(err)
		return
	}
}
