package tools

import (
	"net/smtp"

	"github.com/sirupsen/logrus"

	"github.com/jordan-wright/email"
)

func Send() (err error) {
	e := &email.Email{
		To:      []string{"test@example.com"},
		From:    "Jordan Wright <test@gmail.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		// Headers: textproto.MIMEHeader{},
	}
	err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", "test@gmail.com", "password123", "smtp.gmail.com"))
	if err != nil {
		logrus.Error("email send err : ", err.Error())
	}
	return
}
