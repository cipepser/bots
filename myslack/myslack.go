package myslack

import (
	"fmt"
	"log"

	"github.com/cipepser/bots/util"
	"github.com/nlopes/slack"
)

type Token struct {
	APIToken string `json:"api_token"`
}

func NewRTM() (*slack.RTM, error) {
	t := &Token{}
	err := util.GetToken("./token/slack_gobot.json", t)
	if err != nil {
		return nil, err
	}

	return slack.New(t.APIToken).NewRTM(), nil
}

func run(api *slack.Client, Msg, ch chan string) int {
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-Msg:
			log.Print("message: ", msg)
			// rtm.SendMessage(rtm.NewOutgoingMessage("*************************NEW MESSAGE!!*************************", "C891LB9PF"))
			rtm.SendMessage(rtm.NewOutgoingMessage(msg, "C891LB9PF"))
			// rtm.SendMessage(rtm.NewOutgoingMessage("***************************************************************", "C891LB9PF"))
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("bot start")
				rtm.SendMessage(rtm.NewOutgoingMessage("bot start", "C891LB9PF"))

				// rtm.SendMessage(rtm.NewOutgoingMessage("bot start", ev.Channel))

			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				// fmt.Println()
				fmt.Println("****", ev.Channel, "****")
				ch <- ev.Msg.Text
				// rtm.SendMessage(rtm.NewOutgoingMessage("new message", ev.Channel))

			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return 1

			}
		}
	}
}
