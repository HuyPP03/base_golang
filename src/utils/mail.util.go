package utils

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	gomail "gopkg.in/gomail.v2"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

// SendMail sends an email using the gomail package
func SendMail(toUsers, subject, text, html string) error {
	LoadEnv()

	host := os.Getenv("MAIL_HOST")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	user := os.Getenv("MAIL_USER")
	pass := os.Getenv("MAIL_PASS")
	from := os.Getenv("MAIL_FROM")

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", toUsers)
	m.SetHeader("Subject", subject)
	if text != "" {
		m.SetBody("text/plain", text)
	}
	if html != "" {
		m.SetBody("text/html", html)
	}

	d := gomail.NewDialer(host, port, user, pass)

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
