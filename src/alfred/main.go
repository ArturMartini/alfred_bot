package main

import (
	"alfred/control"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
)

func main() {

	api := slack.New(
		"your_slack_token",
		slack.OptionDebug(true),
		slack.OptionLog(log.New(os.Stdout, "alfred-bot: ", log.Lshortfile|log.LstdFlags)),
	)

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		fmt.Print("Event Received: ")
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			resp := control.ProcessMessage(ev)
			rtm.SendMessage(rtm.NewOutgoingMessage(resp.Text, ev.Channel))
		}
	}
}
