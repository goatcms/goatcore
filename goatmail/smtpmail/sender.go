package smtpmail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"

	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/workers/jobsync"
	"github.com/goatcms/goatcore/workers/wio"
)

type MailSender struct {
	config Config
}

func NewMailSender(config Config) *MailSender {
	return &MailSender{
		config: config,
	}
}

func (ms *MailSender) Send(mail *goatmail.Mail, lc *jobsync.Lifecycle) error {
	if len(mail.To) < 1 {
		return fmt.Errorf("must define one or more recipient")
	}

	auth := smtp.PlainAuth(ms.config.AuthIdentity, ms.config.AuthUsername, ms.config.AuthPassword, ms.config.AuthHost)

	host, _, _ := net.SplitHostPort(ms.config.SmtpAddr)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", ms.config.SmtpAddr, tlsconfig)
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		return err
	}

	if err = c.Auth(auth); err != nil {
		return err
	}

	if err = c.Mail(mail.From.Address); err != nil {
		return err
	}

	if err = c.Rcpt(mail.To[0].Address); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	smtpreader, err := FormatMail(mail, lc)
	if err != nil {
		return err
	}

	wio.Copy(w, smtpreader, lc)
	err = w.Close()
	if err != nil {
		lc.Error(err)
	}
	c.Quit()

	return nil

}
