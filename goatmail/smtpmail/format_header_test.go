package smtpmail

import (
	"io"
	"net/mail"
	"strings"
	"testing"
	"time"

	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/workers"
	"github.com/goatcms/goatcore/workers/jobsync"
)

func TestFormatMail(t *testing.T) {
	t.Parallel()
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

	if len(lifecycle.Errors()) != 0 {
		t.Errorf("%v", lifecycle.Errors())
		return
	}

	m, err := mail.ReadMessage(r)
	if err != nil {
		t.Error(err)
		return
	}

	fromList, err := m.Header.AddressList("From")
	if err != nil {
		t.Error(err)
		return
	}
	if len(fromList) != 1 {
		t.Errorf("incorrect from list length (expected: %v and get: %v", 1, len(fromList))
		return
	}
	if fromList[0].Name != "fromname" {
		t.Errorf("incorrect from name (expected: %v and get: %v", "fromname", fromList[0].Name)
		return
	}
	if fromList[0].Address != "from@goatcms.com" {
		t.Errorf("incorrect from address (expected: %v and get: %v", "from@goatcms.com", fromList[0].Address)
		return
	}

	toList, err := m.Header.AddressList("To")
	if err != nil {
		t.Error(err)
		return
	}
	if toList[0].Name != "toname" {
		t.Errorf("incorrect to name (expected: %v and get: %v", "toname", toList[0].Name)
		return
	}
	if toList[0].Address != "to@goatcms.com" {
		t.Errorf("incorrect to address (expected: %v and get: %v", "to@goatcms.com", toList[0].Address)
		return
	}

	if m.Header.Get("Subject") != "Goatcore subject" {
		t.Errorf("incorrect Subject %v", m.Header.Get("Subject"))
		return
	}

	if m.Header.Get("Subject") != "Goatcore subject" {
		t.Errorf("incorrect Subject %v", m.Header.Get("Subject"))
		return
	}
	if m.Header.Get("Date") != mailtime.Format(time.RFC1123Z)+" (UTC)" {
		t.Errorf("incorrect date %v (expect: %v)", m.Header.Get("Date"), mailtime.Format(time.RFC1123Z))
		return
	}
}
