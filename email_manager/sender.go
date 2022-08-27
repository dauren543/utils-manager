package email_manager

import (
	"bytes"
	"errors"
	"html/template"
	"log"
)

type sendEmailInput struct {
	To      string
	Subject string
	Body    string
}

func (e *sendEmailInput) validate() error {
	if e.To == "" {
		return errors.New("empty to")
	}

	if e.Subject == "" || e.Body == "" {
		return errors.New("empty subject/body")
	}

	if !isEmailValid(e.To) {
		return errors.New("invalid to email")
	}

	return nil
}

func (e *sendEmailInput) generateBodyFromHTML(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		log.Printf("failed to parse file %s:%s", templateFileName, err.Error())
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}

	e.Body = buf.String()

	return nil
}