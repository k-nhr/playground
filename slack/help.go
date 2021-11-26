package main

import "github.com/nlopes/slack"

var (
	commands = map[string]string{
		"help":                "Displays all of the help commands.",
		"random [n]":          "The number of people specified is randomly selected from this channel.",
		"image \"<keyword>\"": "Retrieve the relevant image specified by the keyword.",
	}
)

func (b *Bot) help() []slack.Attachment {
	fields := make([]slack.AttachmentField, 0)

	for k, v := range commands {
		fields = append(fields, slack.AttachmentField{
			Title: "@" + botName + " " + k,
			Value: v,
		})
	}

	attachment := []slack.Attachment{slack.Attachment{
		Pretext: botName + " Command List",
		Color:   "#B733FF",
		Fields:  fields,
	}}
	return attachment
}
