package smtpmail

import (
	"encoding/base64"
	"net/smtp"
	"strings"

	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/varutil"
)

type MailSender struct {
	config Config
}

func NewMailSender(config Config) *MailSender {
	return &MailSender{
		config: config,
	}
}

func (ms *MailSender) Send(mail *goatmail.Mail) error {
	msg, err := ms.FormatMail(mail)
	if err != nil {
		return err
	}
	if err := smtp.SendMail(ms.config.SmtpAddr,
		smtp.PlainAuth(ms.config.AuthIdentity, ms.config.AuthUsername, ms.config.AuthPassword, ms.config.AuthHost),
		mail.From.Address, mail.ToAddrs(), []byte(msg)); err != nil {
		return err
	}
	return nil
}

func (ms *MailSender) FormatMail(mail *goatmail.Mail) (string, error) {
	// from
	protocol := "From: " + mail.From.String() + "\n"
	protocol += "Reply-To: " + mail.From.Address + "\n"
	// to
	to := ""
	for i, recipant := range mail.To {
		if i > 0 {
			to += ", " + recipant.String()
		} else {
			to += recipant.String()
		}
	}
	protocol += "To: " + to + "\n"
	// subject
	protocol += "Subject: " + EscepeSubject(mail.Subject) + "\n"
	// content
	boundary := varutil.RandString(20, varutil.AlphaNumericBytes)
	protocol += "Content-Type: multipart/alternative; boundary=\"" + boundary + "\"\n"
	for mime, content := range mail.Body {
		protocol += "\n--" + boundary + "\n"
		protocol += "Content-Type: " + mime + "; charset=\"utf-8\"\n"
		protocol += "Content-Transfer-Encoding: base64\n\n"
		protocol += base64.StdEncoding.EncodeToString([]byte(content))
		protocol += boundary + "\n"
	}
	protocol += "\n\n--" + boundary + "--\n"
	return protocol, nil
}

func EscepeSubject(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	return s
}
