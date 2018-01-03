package main

import (
	"log"

	"github.com/nlopes/slack"

	"./myslack"
)

func main() {
	rtm, err := myslack.NewRTM()
	if err != nil {
		panic(err)
	}
	go rtm.ManageConnection()

	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.HelloEvent:
				log.Print("bot start")
			case *slack.MessageEvent:
				log.Printf("Message: %v\n", ev)
				rtm.SendMessage(rtm.NewOutgoingMessage("new message", ev.Channel))
			case *slack.InvalidAuthEvent:
				log.Print("Invalid credentials")
				return
			}
		}
	}
}
