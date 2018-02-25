package goatmail

import (
	"io"
	"net/mail"
	"time"
)

// Address is object represent a single email address
type Address mail.Address

// String convert a address to smtp representation like "Sebastian Pożoga <sebastian@goatcms.com>"
func (a Address) String() string {
	if a.Name == "" {
		return a.Address
	}
	return "\"" + a.Name + "\" <" + a.Address + ">"
}

// Attachment is object represent a single email attachment
type Attachment struct {
	Name   string
	MIME   string
	Reader io.Reader
}

// Mail is object represent a single email
type Mail struct {
	Date        time.Time
	From        Address
	To          []Address
	Subject     string
	Body        map[string]io.Reader
	Attachments []Attachment
}

// ToAddrs repare a string list cntains all recipients of the e-mail
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
