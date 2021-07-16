package ds

import (
	"github.com/nlopes/slack"
)

var (
	BotID   string
	BotName string
)

type Bot struct {
	api *slack.Client
	rtm *slack.RTM
}

func NewBot(token string) *Bot {
	bot := new(Bot)
	bot.api = slack.New(token)
	bot.rtm = bot.api.NewRTM()
	return bot
}
