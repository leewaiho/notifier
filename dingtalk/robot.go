package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/LeeWaiHo/notifier/crypto"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Robot struct {
	accessToken string
	baseURL     string
	secretKey   string
	hc          *http.Client
}

type Option struct {
	AccessToken    string
	SecretKey      string
	RequestTimeout *time.Duration
}

func NewRobot(option *Option) *Robot {
	r := new(Robot)
	r.baseURL = fmt.Sprint(WebHookURLPrefix, "?access_token=", option.AccessToken)
	r.secretKey = option.SecretKey
	timeoutVal := DefaultRequestTimeout
	if option.RequestTimeout != nil {
		timeoutVal = *option.RequestTimeout
	}
	r.hc = &http.Client{
		Timeout: timeoutVal,
	}
	return r
}

func (r *Robot) Send(data interface{}) (string, error) {
	req, err := r.buildMessageRequest(data)
	if err != nil {
		return "", nil
	}
	resp, e := r.hc.Do(req)
	if e != nil {
		return "", e
	}
	defer resp.Body.Close()
	respBytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", e
	}
	return string(respBytes), nil
}

func (r *Robot) buildMessageRequest(messageData interface{}) (*http.Request, error) {
	bs, e := json.Marshal(messageData)
	if e != nil {
		return nil, e
	}
	targetURL, _ := url.Parse(r.baseURL)
	if len(r.secretKey) > 0 {
		milliTimestamp := fmt.Sprint(time.Now().UnixNano() / 1000000)
		q := targetURL.Query()
		q.Set("timestamp", milliTimestamp)
		q.Set("sign", makeSignature(milliTimestamp, r.secretKey))
		targetURL.RawQuery = q.Encode()
	}
	req, e := http.NewRequest("POST", targetURL.String(), bytes.NewReader(bs))
	if e != nil {
		return nil, e
	}
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	return req, nil
}

func makeSignature(milliTimestamp, secretKey string) string {
	contentBytes := []byte(milliTimestamp + "\n" + secretKey)
	return crypto.Base64(crypto.HS256(contentBytes, []byte(secretKey)))
}
