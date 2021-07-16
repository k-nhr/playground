package main

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/nlopes/slack"
)

var login *Login

func (b *Bot) punchClock() ([]slack.Attachment, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Jar: jar, Transport: tr}

	var reqBody []byte
	reqBody = []byte(`company=1000&username=` + login.ID + `&password=` + login.PW + `&loginClickFlg=true`)
	req, err := http.NewRequest(http.MethodPost, "https://192.168.241.10/logtime/login.do", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		return nil, fmt.Errorf("StatusCode=%d", res.StatusCode)
	}

	setCookieUrl, err := url.Parse("http://localhost/set_cookie")
	if err != nil {
		return nil, err
	}
	cookies := jar.Cookies(setCookieUrl)
	fmt.Printf("%v\n", cookies)

	reqBody = []byte(`actionPattern=NewExe&dakokuCallSb=topmenu&dakokuKubun=$1`)
	req, err = http.NewRequest(http.MethodPost, "https://192.168.241.10/logtime/webdakoku.do", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res.Body.Close()
	res, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 201 {
		return nil, fmt.Errorf("StatusCode=%d", res.StatusCode)
	}

	attachment := []slack.Attachment{slack.Attachment{
		Text: "打刻完了 :muscle:",
	}}

	return attachment, nil
}
