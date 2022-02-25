package mail

import (
	"bytes"
	"crypto/tls"
	"log"
	"os"
	"text/template"

	gomail "gopkg.in/mail.v2"
)

func Send(name string, email string, token string) {
	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASS")
	smtpHost := os.Getenv("SMTP_HOST")

	mail := gomail.NewMessage()

	// Set E-Mail sender
	mail.SetHeader("From", from)

	// Set E-Mail receivers
	mail.SetHeader("To", email)

	// Set E-Mail subject
	mail.SetHeader("Subject", "Recuperação de senha")

	temp, err := template.ParseFiles(os.Getenv("GO_PATH") + "/users-ms/adapter/mail/template/recover-password.html")

	if err != nil {
		log.Fatal(err)
		return
	}

	// Set E-Mail body. You can set plain text or html with text/html
	var tpl bytes.Buffer
	if err := temp.Execute(&tpl, struct {
		Name  string
		Token string
	}{
		Name:  name,
		Token: token,
	}); err != nil {
		log.Fatal(err)
		return
	}

	mail.SetBody("text/html", tpl.String())

	// Settings for SMTP server
	d := gomail.NewDialer(smtpHost, 587, from, password)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(mail); err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Email Sent Successfully for: " + email)
}
