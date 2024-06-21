package notification

import "testing"

func TestHWSMSSend(t *testing.T) {
	params := &HWSMSSendReq{
		ApiAddress:    "https://smsapi.cn-north-4.myhuaweicloud.com:443",
		AppKey:        "w9C2loxU75mxxxxxxx",
		AppSecret:     "N1L4GTt6wxxxxxx",
		Sender:        "1069368---------",
		TemplateId:    "4e66940a--------",
		TemplateParas: "[\"313246\"]",
		Receiver:      "+8618117------",
	}
	err := HWSMSSend(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("发送短信成功")
}
