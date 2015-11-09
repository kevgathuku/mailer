// Adapted from the Google App Engine github.com/scorredoira/email packages.
package mail

import (
	"errors"
	"net/mail"
	"net/smtp"

	"github.com/kevgathuku/mailer/Godeps/_workspace/src/gopkg.in/jordan-wright/email.v1"
)

// Mailer encapsulates data used for sending email.
type Mailer struct {
	Auth    smtp.Auth
	Address string
}

// NewMailer creates a new Mailer.
func NewMailer(username, password, host, port string) Mailer {
	return Mailer{
		Auth: smtp.PlainAuth(
			"",
			username,
			password,
			host,
		),
		Address: host + ":" + port,
	}
}

// NewMessage builds a new message instance
func NewMessage(from, to *mail.Address, subject string,
	htmlContent string) (*email.Email, error) {
	msg := &email.Email{
		From:    from.String(),
		To:      []string{to.String()},
		Subject: subject,
		HTML:    []byte(htmlContent),
	}

	return msg, nil
}

// Send sends an email Message.
func (m *Mailer) Send(msg *email.Email) error {
	err := msg.Send(
		m.Address,
		m.Auth,
	)
	if err != nil {
		return errors.New("Error sending email: " + err.Error())
	}

	return nil
}
