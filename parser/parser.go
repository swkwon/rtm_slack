package parser

import (
	"strings"
	"test/rtm_slack/def"

	"github.com/nlopes/slack"
)

func sendMessage(rtm *slack.RTM, msg string) {
	rtm.SendMessage(rtm.NewOutgoingMessage(msg, def.ChannelID))
}

func ParseMessage(rtm *slack.RTM, message string) {
	if strings.HasPrefix(message, "윌리엄") == true {
		sendMessage(rtm, "Hello!")
		// rtm.SendMessage(rtm.NewOutgoingMessage("Yes sir!", def.ChannelID))
	}
}
