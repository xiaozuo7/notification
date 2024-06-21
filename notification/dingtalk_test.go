package notification

import "testing"

func TestDingTalkSend(t *testing.T) {
	params := &DingTalkSendReq{
		WebHook:   "https://oapi.dingtalk.com/robot/send?access_token=xxxxxx",
		Sign:      "SECfxxxxxxxxxxxxxxxxxxxxf2exxxxx",
		Content:   "我是小爱同学，Are u ok?",
		IsAtAll:   true,
		AtMobiles: []string{},
	}
	err := DingTalkSend(params)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("发送消息成功")
}
