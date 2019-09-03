package smtpmail

import (
	"crypto/tls"
	"io"
	"net"
	"net/smtp"

	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/varutil/goaterr"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// MailSender provide dial send stream api
type MailSender struct {
	config Config
}

// NewMailSender create new MailSender isntance
func NewMailSender(config Config) *MailSender {
	return &MailSender{
		config: config,
	}
}

// Send transmit email to server. Use Lifecycle for communication streams.
func (ms *MailSender) Send(mail *goatmail.Mail, lc *jobsync.Lifecycle) error {
	if len(mail.To) < 1 {
		return goaterr.Errorf("must define one or more recipient")
	}

	host, _, _ := net.SplitHostPort(ms.config.SMTPAddr)
	auth := smtp.PlainAuth(ms.config.AuthIdentity, ms.config.AuthUsername, ms.config.AuthPassword, host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", ms.config.SMTPAddr, tlsconfig)
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

	io.Copy(w, smtpreader)
	err = w.Close()
	if err != nil {
		lc.Error(err)
	}
	c.Quit()

	return nil

}
