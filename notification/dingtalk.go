package notification

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type DingTalkSendReq struct {
	WebHook   string   `json:"webhook"`   // 机器人webhook地址
	Sign      string   `json:"sign"`      // 签名
	Content   string   `json:"content"`   // 内容
	IsAtAll   bool     `json:"isAtAll"`   // 是否艾特所有人
	AtMobiles []string `json:"atMobiles"` // 被@的群成员手机号
}

// DingTalkSend 钉钉机器人消息发送
func DingTalkSend(params *DingTalkSendReq) error {
	url := params.WebHook
	// 如果开启了验签，需要进行加签
	if params.Sign != "" {
		timestamp := fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
		encryptSign := generateSign(timestamp, params.Sign)
		url = fmt.Sprintf("%v&timestamp=%v&sign=%v", params.WebHook, timestamp, encryptSign)
	}

	message := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": params.Content,
		},
		"at": map[string]interface{}{
			"isAtAll":   params.IsAtAll,
			"atMobiles": params.AtMobiles,
		},
	}

	jsonValue, _ := json.Marshal(message)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
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

func generateSign(timestamp, secret string) string {
	secretEnc := []byte(secret)
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	stringToSignEnc := []byte(stringToSign)

	h := hmac.New(sha256.New, secretEnc)
	h.Write(stringToSignEnc)
	hmacCode := h.Sum(nil)

	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(hmacCode))
	return sign
}
