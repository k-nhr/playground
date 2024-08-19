package main

import (
	"time"

	"github.com/nlopes/slack"
	"github.com/sclevine/agouti"
)

func (b *Bot) punchClock() ([]slack.Attachment, error) {

	// ブラウザはChromeを指定して起動
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		return nil, err
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		return nil, err
	}
	if err := page.Navigate("https://qiita.com/"); err != nil {
		return nil, err
	}
	if err := page.Screenshot("c:\\Screenshot01.png"); err != nil {
		return nil, err
	}
	/*
		// ログインページに遷移
		if err := page.Navigate("https://ltap.sis.saison.co.jp/logtime/"); err != nil {
		return nil, err
		}
		time.Sleep(1 * time.Second)

		// ID, Passの要素を取得し、値を設定
		if err := page.FindByID("username").Fill("1011626"); err != nil {
		return nil, err
		}
		time.Sleep(1 * time.Second)

		if err := page.FindByID("password").Fill("lt1011626"); err != nil {
		return nil, err
		}
		time.Sleep(1 * time.Second)

		// formをサブミット
		if err := page.FindByID("loginClickFlg").Fill("true"); err != nil {
		return nil, err
		}
		if err := page.FindByID("pleaseWait").Fill("block"); err != nil {
		return nil, err
		}
		if err := page.FindByID("loginButton").Fill("none"); err != nil {
		return nil, err
		}
		time.Sleep(1 * time.Second)

		if err := page.FindByID("loginForm").Submit(); err != nil {
		return nil, err
		}
	*/
	// 処理完了後、3秒間ブラウザを表示しておく
	time.Sleep(3 * time.Second)

	attachment := []slack.Attachment{slack.Attachment{
		Text: "打刻完了 :muscle:",
	}}

	return attachment, nil
}
