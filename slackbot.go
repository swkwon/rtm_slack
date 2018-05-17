package main

import (
	"fmt"
	"log"
	"os"

	"test/rtm_slack/def"
	"test/rtm_slack/parser"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(def.APIToken)
	logger := log.New(os.Stdout, "slack-bot: ", log.Lshortfile|log.LstdFlags)
	slack.SetLogger(logger)
	api.SetDebug(false)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.HelloEvent:
			// Ignore hello

		case *slack.ConnectedEvent:
			fmt.Println("Infos:", ev.Info)
			fmt.Println("Connection counter:", ev.ConnectionCount)
			// Replace C2147483705 with your Channel ID
			// rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", "CAQMGC35X"))

		case *slack.MessageEvent:
			fmt.Printf("Message: %s\n", ev.Text)
			parser.ParseMessage(rtm, ev.Text)
		// case *slack.PresenceChangeEvent:
		// 	fmt.Printf("Presence Change: %v\n", ev)

		// case *slack.LatencyReport:
		// 	fmt.Printf("Current latency: %v\n", ev.Value)

		// case *slack.RTMError:
		// 	fmt.Printf("Error: %s\n", ev.Error())

		// case *slack.InvalidAuthEvent:
		// 	fmt.Printf("Invalid credentials")
		// 	return

		default:
		}
	}
}
