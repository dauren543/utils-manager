package email_manager

import (
	"errors"
	"fmt"
	"github.com/go-gomail/gomail"
)

type emailManager struct {
	from string
	pass string
	host string
	port int
}

func NewEmailManager(from, pass, host string, port int) (*emailManager, error) {
	if !isEmailValid(from) {
		return nil, errors.New("invalid from email")
	}
	return &emailManager{from: from, pass: pass, host: host, port: port}, nil
}

func (s *emailManager) send(input sendEmailInput) error {
	if err := input.validate(); err != nil {
		return err
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", input.To)
	msg.SetHeader("Subject", input.Subject)
	msg.SetBody("text/html", input.Body)
	dialer := gomail.NewDialer(s.host, s.port, s.from, s.pass)
	if err := dialer.DialAndSend(msg); err != nil {
		return errors.New(fmt.Sprintf("failed to sent email via smtp, error: %v", err))
	}

	return nil
}

func (s *emailManager) SendEmail(to, subject string, templatePath string, body interface{}) error {
	em := sendEmailInput{
		To:      to,
		Subject: subject,
	}

	err := em.generateBodyFromHTML(templatePath, body)
	if err != nil {
		return err
	}

	return s.send(em)
}
