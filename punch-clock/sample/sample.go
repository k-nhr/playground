package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func main() {
	os.Setenv("NO_PROXY", "*.sis.saison.co.jp")

	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Jar: jar, Transport: tr}

	/*
		var reqBody []byte
		reqBody = []byte(`company=1000&username=1011626&password=lt1011626&loginClickFlg=true`)
		req, err := http.NewRequest(http.MethodPost, "https://ltap.sis.saison.co.jp/logtime/login.do", bytes.NewBuffer(reqBody))
	*/
	form := url.Values{}
	form.Add("company", "1000")
	form.Add("username", "1011626")
	form.Add("password", "lt1011626")
	form.Add("loginClickFlg", "true")
	req, err := http.NewRequest(http.MethodPost, "https://ltap.sis.saison.co.jp/logtime/login.do", strings.NewReader(form.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Errorf("StatusCode=%d", res.StatusCode))
	}

	setCookieURL, err := url.Parse("https://ltap.sis.saison.co.jp/logtime/login.do")
	if err != nil {
		panic(err)
	}
	cookies := jar.Cookies(setCookieURL)
	fmt.Printf("%v\n", cookies)

	/*
		reqBody = []byte(`actionPattern=NewExe&dakokuCallSb=topmenu&dakokuKubun=01`)
		req, err = http.NewRequest(http.MethodPost, "https://ltap.sis.saison.co.jp/logtime/webdakoku.do", bytes.NewBuffer(reqBody))
	*/
	form2 := url.Values{}
	form2.Add("actionPattern", "NewExe")
	form2.Add("dakokuCallSb", "topmenu")
	form2.Add("dakokuKubun", "01")
	req, err = http.NewRequest(http.MethodPost, "https://ltap.sis.saison.co.jp/logtime/webdakoku.do", strings.NewReader(form2.Encode()))
	if err != nil {
		panic(err)
	}

	res.Body.Close()
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic(fmt.Errorf("StatusCode=%d", res.StatusCode))
	}

	return
}
