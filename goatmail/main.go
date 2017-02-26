package goatmail

import (
	"io"
	"net/mail"
	"time"
)

type Address mail.Address

func (a Address) String() string {
	if a.Name == "" {
		return a.Address
	}
	return "\"" + a.Name + "\" <" + a.Address + ">"
}

type Attachment struct {
	Name   string
	MIME   string
	Reader io.Reader
}

type Mail struct {
	Date        time.Time
	From        Address
	To          []Address
	Subject     string
	Body        map[string]string
	Attachments []Attachment
}

func (mail Mail) ToAddrs() []string {
	arr := make([]string, len(mail.To))
	for i, to := range mail.To {
		arr[i] = to.Address
	}
	return arr
}

type MailSender interface {
	Send(*Mail)
}
