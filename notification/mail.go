package notification

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type MailSendReq struct {
	From        string   `json:"from"`        // 发件人
	To          []string `json:"to"`          // 收件人
	User        string   `json:"user"`        // smtp用户名
	Password    string   `json:"password"`    // smtp密码/授权码
	Host        string   `json:"host"`        // smtp host  [465, 587]
	Port        int      `json:"port"`        // smtp port
	Subject     string   `json:"subject"`     // 邮件主题
	ContentType string   `json:"contentType"` // 邮件内容类型 text/plain text/html
	Content     string   `json:"content"`     // 邮件内容
}

// MailService 发送邮件
func MailSend(req *MailSendReq) error {
	m := gomail.NewMessage()
	m.SetHeader("From", req.From)
	m.SetHeader("To", req.To...)
	m.SetHeader("Subject", req.Subject)
	m.SetBody(req.ContentType, req.Content)
	d := gomail.NewDialer(req.Host, req.Port, req.User, req.Password)

	// 发送邮件
	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("邮件发送失败，error: %v", err)
	}
	return nil
}
