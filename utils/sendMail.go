package utils

import (
	"net/smtp"
	"os"
)

func SendEmail(to, subject, body string) error {
	from := os.Getenv("SMTP_USERNAME")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body + "\r\n")

	addr := smtpHost + ":" + smtpPort
	return smtp.SendMail(addr, auth, from, []string{to}, msg)
}
