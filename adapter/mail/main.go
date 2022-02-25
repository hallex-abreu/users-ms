package mail

import (
	"crypto/tls"
	"log"
	"os"

	gomail "gopkg.in/mail.v2"
)

func Send(email string, token string) {
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

	// Set E-Mail body. You can set plain text or html with text/html
	mail.SetBody("text/html", "<h1>Esqueceu sua senha?</h1></br><a href='teste/"+token+"'>Clique aqui para recuperar sua senha!</a>")

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

	log.Println("Email Sent Successfully!")
}
