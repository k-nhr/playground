package main

import (
	"github.com/nlopes/slack"
)

type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

func newBot(token string) *Bot {
	bot := new(Bot)
	bot.api = slack.New(token)
	bot.api.SetDebug(true)
	bot.rtm = bot.api.NewRTM()
	return bot
}
