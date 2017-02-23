package goatmail

import "net/mail"

type Address mail.Address

func (a Address) String() string {
	if a.Name == "" {
		return a.Address
	}
	return "\"" + a.Name + "\" <" + a.Address + ">"
}

type Mail struct {
	From    Address
	To      []Address
	Subject string
	Body    map[string]string
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
