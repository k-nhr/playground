package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nlopes/slack"
)

const botIcon = ":hubot:"

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

func (b *Bot) handleResponse(user, text, channel string) {
	commandArray := strings.Fields(text)

	var cmd string
	l := len(commandArray)
	if l <= 1 {
		cmd = "help"
	} else {
		cmd = commandArray[1]
	}

	var attachments []slack.Attachment
	var err error
	switch cmd {
	case "random":
		num := 1
		if l >= 3 {
			num, err = strconv.Atoi(commandArray[2])
			if err != nil {
				break
			}
		}
		attachments, err = b.random(channel, num)
	case "image":
		if l != 3 {
			err = fmt.Errorf("argument is missing")
			break
		}
		attachments, err = b.image(commandArray[2])
	case "help":
		attachments = b.help()
	case "kill":
		b.rtm.Disconnect()
		return
	default:
		attachments = b.help()
	}

	if err != nil {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(fmt.Sprintf("Sorry %s is error... %s", cmd, err), channel))
		return
	}

	params := slack.PostMessageParameters{
		Attachments: attachments,
		Username:    botName,
	}

	_, _, err = b.api.PostMessage(channel, "", params)
	if err != nil {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(fmt.Sprintf("Sorry %s is error... %s", cmd, err), channel))
		return
	}
}
