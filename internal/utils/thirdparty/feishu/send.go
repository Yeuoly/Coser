package email

import "gopkg.in/gomail.v2"

type BocchiEmailManager struct {
	host string
	port int
	user string
	pass string
}

func NewBocchiEmailManager(host string, port int, user string, pass string) *BocchiEmailManager {
	return &BocchiEmailManager{
		host: host,
		port: port,
		user: user,
		pass: pass,
	}
}

func (c *BocchiEmailManager) SendMail(to string, subject string, text string) error {
	d := gomail.NewDialer(c.host, c.port, c.user, c.pass)
	m := gomail.NewMessage()
	m.SetHeader("From", c.user)
	m.SetHeader("To", to)
	m.SetAddressHeader("Miduoduo", c.user, "Miduoduo")
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", text)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
