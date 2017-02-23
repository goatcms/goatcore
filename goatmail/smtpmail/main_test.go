package smtpmail

import (
	"testing"

	"github.com/goatcms/goatcore/goatmail"
)

func TestSendEmail(t *testing.T) {
	config, err := LoadTestConfig()
	if err != nil {
		t.Error(err)
		return
	}

	ms := NewMailSender(config.SenderConfig)

	mail := &goatmail.Mail{
		From: goatmail.Address{
			Name:    config.FromAddress,
			Address: config.FromAddress,
		},
		To: []goatmail.Address{goatmail.Address{
			Name:    config.ToAddress,
			Address: config.ToAddress,
		}},
		Subject: "Goatcore subject",
		Body: map[string]string{
			"text/plain": "some content",
			"text/html":  "some <b>content</b>",
		},
	}

	if err := ms.Send(mail); err != nil {
		t.Error(err)
		return
	}
}
