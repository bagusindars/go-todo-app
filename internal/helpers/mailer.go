package helpers

import (
	"log"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, message string) {
	host := os.Getenv("SMTP_HOST")
	port, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))
	user := os.Getenv("SMTP_USER")
	pass := os.Getenv("SMTP_PASS")

	mail := gomail.NewMessage()
	mail.SetHeader("From", user)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", message)

	dialer := gomail.NewDialer(host, port, user, pass)

	if err := dialer.DialAndSend(mail); err != nil {
		log.Printf("Failed to send email to %s : %s", to, err.Error())
		return
	}

	log.Printf("Email sent")
}
