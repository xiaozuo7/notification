package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type QYWinXinSendReq struct {
	WebHook             string   `json:"webhook"`               // 机器人webhook地址
	Content             string   `json:"content"`               // 内容
	MentionedMobileList []string `json:"mentioned_mobile_list"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
}

// QYWinXinSend 企业微信机器人消息发送
func QYWinXinSend(params *QYWinXinSendReq) error {
	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content":               params.Content,
			"mentioned_mobile_list": params.MentionedMobileList,
		},
	}

	jsonValue, err := json.Marshal(message)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", params.WebHook, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	type response struct {
		ErrCode json.Number `json:"errcode"`
		ErrMsg  string      `json:"errmsg"`
	}
	var res response
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return err
	}
	if res.ErrCode != "0" {
		return fmt.Errorf("返回码：%v, 错误原因：%v", res.ErrCode, res.ErrMsg)
	}

	return nil
}
