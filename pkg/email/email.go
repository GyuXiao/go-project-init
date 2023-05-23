package email

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

// SMTP: Simple Mail Transfer Protocol

type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

type Email struct {
	*SMTPInfo
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

func (e *Email) SendEmail(to []string, subject, body string) error {
	// 先创建一个消息实例，然后设置邮件的一些必要信息，分别是：
	// 发件人，收件人，邮件主题，邮件内容
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	// 创建一个新的 SMTP 拨号实例，设置对应的拨号信息用于连接 SMTP 服务器
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	// 调用 DialAndSend 打开与 SMTP 服务器的连接并发送 email
	return dialer.DialAndSend(m)
}
