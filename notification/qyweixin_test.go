package notification

import "testing"

func TestQYWinXinSend(t *testing.T) {
	params := &QYWinXinSendReq{
		WebHook:             "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=a492xxxxxxxxxxxxx96",
		Content:             "我是小爱同学，Are u ok?",
		MentionedMobileList: []string{"@all"},
	}
	err := QYWinXinSend(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("发送消息成功")
}
