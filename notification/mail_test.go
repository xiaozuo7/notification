package notification

import "testing"

func TestMail(t *testing.T) {
	params := &MailSendReq{
		From:        "8xxxxxxx9@qq.com",
		To:          []string{"chxxxng@qq.com"},
		User:        "89xxxxxx9@qq.com",
		Password:    "frxxxxxxxxxba",
		Host:        "smtp.qq.com",
		Port:        465,
		Subject:     "测试邮件",
		ContentType: "text/plain",
		Content:     "内容11111111",
	}
	err := MailSend(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("发送邮件成功")
}
