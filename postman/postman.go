package postman

import "gopkg.in/gomail.v2"

type Channel interface {
	Send(mail *Mail) error
}

type Mail struct {
	Receivers    []string
	Subject      string
	Content      string
	AttachedFile []string
}

func NewMail(receivers []string, subject, content string) *Mail {
	return &Mail{
		Receivers: receivers,
		Subject:   subject,
		Content:   content,
	}
}

func (m *Mail) Attach(files ...string) *Mail {
	m.AttachedFile = append(m.AttachedFile, files...)
	return m
}

func (m *Mail) toMessage() *gomail.Message {
	return gomail.NewMessage(func(msg *gomail.Message) {
		msg.SetHeader("Subject", m.Subject)
		msg.SetHeader("To", m.Receivers...)
		msg.SetBody("text/html", m.Content)
	})
}

type Credential struct {
	ServerHost string
	ServerPort int
	Username   string
	Password   string
	Nickname   string
}

type Postman struct {
	conn *Credential
}

func (p *Postman) newDialer() *gomail.Dialer {
	return gomail.NewDialer(p.conn.ServerHost, p.conn.ServerPort, p.conn.Username, p.conn.Password)
}

func (p *Postman) Send(mail *Mail) error {
	return p.newDialer().DialAndSend(mail.toMessage())
}
