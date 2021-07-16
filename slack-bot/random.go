package main

import (
	"math/rand"
	"time"

	"github.com/nlopes/slack"
)

func (b *Bot) random(channel string, num int) ([]slack.Attachment, error) {
	var attachments []slack.Attachment
	var members []string

	if c, err := b.api.GetChannelInfo(channel); err != nil {
		g, err := b.api.GetGroupInfo(channel)
		if err != nil {
			return attachments, err
		}
		members = g.Members
	} else {
		members = c.Members
	}

	if num > len(members)-1 {
		num = len(members) - 1
	}

	var i int
	var selected []int
	var attachment slack.Attachment

	rand.Seed(time.Now().UnixNano())
	for {
		i = rand.Intn(len(members))
		if members[i] != botID && !contains(selected, i) {
			attachment.Pretext = "選ばれたのは <@" + members[i] + "> でした :tea:"
			attachment.Color = "#B733FF"
			attachments = append(attachments, attachment)
			selected = append(selected, i)
		}
		if len(selected) == num {
			break
		}
	}
	return attachments, nil
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
