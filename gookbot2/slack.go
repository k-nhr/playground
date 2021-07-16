package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/nlopes/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

type SlackListener struct {
	client    *slack.Client
	botID     string
	channelID string
}

// LstenAndResponse listens slack events and response
// particular messages. It replies by slack message button.
func (s *SlackListener) ListenAndResponse() {
	rtm := s.client.NewRTM()

	// Start listening slack events
	go rtm.ManageConnection()

	// Handle slack events
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			if err := s.handleMessageEvent(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}
		}
	}
}

// handleMesageEvent handles message events.
func (s *SlackListener) handleMessageEvent(ev *slack.MessageEvent) error {
	// Only response in specific channel. Ignore else.
	if ev.Channel != s.channelID {
		log.Printf("%s %s", ev.Channel, ev.Msg.Text)
		return nil
	}

	// Only response mention to bot. Ignore else.
	if !strings.HasPrefix(ev.Msg.Text, fmt.Sprintf("<@%s> ", s.botID)) {
		return nil
	}

	// Parse message
	m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	if len(m) == 0 || m[0] != "asakai" || m[0] != "yukai" {
		return fmt.Errorf("invalid message")
	}

	var attachments []slack.Attachment
	var err error
	switch m[0] {
	case "asakai":
		attachments, err = b.asakai(channel, num)
	case "yukai":
		attachments, err = b.yukai(m[2])
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

	if _, _, err := s.client.PostMessage(ev.Channel, "", params); err != nil {
		b.rtm.SendMessage(b.rtm.NewOutgoingMessage(fmt.Sprintf("Sorry %s is error... %s", cmd, err), channel))
		return
	}

	return nil
}

func (bot *Bot) asakai(members []member) {
	ticker := time.NewTicker(20 * time.Hour)

	for i, m := range members {
		bot.yesterdaysTask(m.Id)
		members[i].status = YESTERDAY
	}

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				exist, i := contains(members, ev.Channel)
				if !exist {
					break
				}
				if ev.Type == "message" && ev.Text != "" {
					setMessage(&members[i], ev.Text)
					bot.reply(&members[i], ev.Channel)
				}

			case *slack.DisconnectedEvent:
				return
			}
		case <-ticker.C:
			return
		}
	}
}

func contains(s []member, e string) (bool, int) {
	for i, v := range s {
		if e == v.Id {
			return true, i
		}
	}
	return false, 0
}

func setMessage(m *member, msg string) {
	if msg == "no" || msg == "-" {
		m.status++
		return
	}
	switch m.status {
	case YESTERDAY:
		m.ymsg = msg
	case TODAY:
		m.tmsg = msg
	case PROBREM:
		m.pmsg = msg
	case HOLIDAY:
		m.hmsg = msg
	}
	m.status++
}
