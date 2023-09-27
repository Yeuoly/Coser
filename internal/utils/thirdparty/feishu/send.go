package email

import "gopkg.in/gomail.v2"

type BillboardsEmailManager struct {
	host string
	port int
	user string
	pass string
}

func NewBillboardsEmailManager(host string, port int, user string, pass string) *BillboardsEmailManager {
	return &BillboardsEmailManager{
		host: host,
		port: port,
		user: user,
		pass: pass,
	}
}

func (c *BillboardsEmailManager) SendMail(to string, subject string, text string) error {
	d := gomail.NewDialer(c.host, c.port, c.user, c.pass)
	m := gomail.NewMessage()
	m.SetHeader("From", c.user)
	m.SetHeader("To", to)
	m.SetAddressHeader("Billboards", c.user, "Billboards")
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", text)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
