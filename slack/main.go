package main

import (
	"os"
	"strings"

	"github.com/nlopes/slack"
)

var (
	botID   string
	botName string
)

func main() {
	token := os.Getenv("SLACK_API_TOKEN")
	bot := NewBot(token)

	go bot.rtm.ManageConnection()

	done := make(chan struct{})
	go func() {
		defer close(done)

		for msg := range bot.rtm.IncomingEvents {
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				botID = ev.Info.User.ID
				botName = ev.Info.User.Name

			case *slack.MessageEvent:
				user := ev.User
				text := ev.Text
				channel := ev.Channel
				if ev.Type == "message" && strings.HasPrefix(text, "<@"+botID+">") {
					bot.handleResponse(user, text, channel)
				}
			case *slack.DisconnectedEvent:
				return
			}
		}
	}()
	<-done
	return
}
