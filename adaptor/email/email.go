package email

import "gopkg.in/gomail.v2"

func (e emailService) SendEmail(to string, subject string, body string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", e.Address)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(e.SmtpHost, e.SmtpPort, e.Address, e.Password)
	return d.DialAndSend(m)
}
