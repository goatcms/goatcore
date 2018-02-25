package smtpmail

import (
	"encoding/base64"
	"io"
	"mime/quotedprintable"
	"strings"
	"time"

	"github.com/goatcms/goatcore/goatmail"
	"github.com/goatcms/goatcore/varutil"
	"github.com/goatcms/goatcore/workers/jobsync"
)

// FormatMail prepare SMTP message
func FormatMail(mail *goatmail.Mail, lc *jobsync.Lifecycle) (io.Reader, error) {
	boundary := varutil.RandString(20, varutil.AlphaNumericBytes)
	header := ""

	// from
	header += "From: " + mail.From.String() + "\n"
	header += "Reply-To: " + mail.From.Address + "\n"
	// to
	to := ""
	for i, recipant := range mail.To {
		if i > 0 {
			to += ", " + recipant.String()
		} else {
			to += recipant.String()
		}
	}
	header += "To: " + to + "\n"
	header += "Subject: " + EscepeSubject(mail.Subject) + "\n"
	header += "MIME-version: 1.0\n"
	header += "Content-Type: multipart/mixed; boundary=\"" + boundary + "\"\n"
	header += "Date: " + mail.Date.Format(time.RFC1123Z) + " (UTC)\n\n"

	// content
	header += "\n--" + boundary + "\n"
	header += "Content-Type: multipart/alternative; boundary=\"content_" + boundary + "\""

	readers := []io.Reader{
		strings.NewReader(header),
	}

	for mime, body := range mail.Body {
		bodyHeader := "\n\n--content_" + boundary + "\n"
		bodyHeader += "Content-Type: " + mime + "; charset=\"utf-8\"\n"
		bodyHeader += "Content-Transfer-Encoding: QUOTED-PRINTABLE\n\n"
		reader, writer := io.Pipe()
		w := quotedprintable.NewWriter(writer)
		go func(bodyReader io.Reader) {
			_, err := io.Copy(w, bodyReader)
			if err != nil {
				lc.Error(err)
			}
			w.Close()
			writer.Close()
		}(body)
		readers = append(readers, strings.NewReader(bodyHeader), reader)
	}
	readers = append(readers, strings.NewReader("\n\n--content_"+boundary+"--\n"))

	//	attachments
	for _, attachment := range mail.Attachments {
		attStr := "\n\n--" + boundary + "\n"
		attStr += "Content-Type: " + attachment.MIME + "; name=\"" + attachment.Name + "\"\n"
		attStr += "Content-Transfer-Encoding: base64\n"
		attStr += "Content-Disposition: attachment; filename=\"" + attachment.Name + "\"\n\n"
		reader, writer := io.Pipe()
		encoder := base64.NewEncoder(base64.StdEncoding, writer)
		go func(reader io.Reader) {
			_, err := io.Copy(encoder, reader)
			if err != nil {
				lc.Error(err)
			}
			encoder.Close()
			writer.Close()
		}(attachment.Reader)
		readers = append(readers, strings.NewReader(attStr), reader)
	}
	readers = append(readers, strings.NewReader("\n\n--"+boundary+"--\n"))

	return io.MultiReader(readers...), nil
}

// EscepeSubject remove incorrect characters
func EscepeSubject(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	return s
}
