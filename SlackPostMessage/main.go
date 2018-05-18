package main

import (
	"fmt"
	"test/rtm_slack/def"

	"github.com/nlopes/slack"
)

func main() {
	api := slack.New(def.APIToken)
	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Color:      "#2eb886",
		AuthorName: "william",
		Pretext:    "[TV]나의 아저씨가 남긴것..",
		Text:       "나의 아저씨 결말까지...",
		Footer:     "TV리포트",
		Title:      "나의 아저씨 기사",
		TitleLink:  "http://www.tvreport.co.kr/?c=news&m=newsview&idx=1057739",
	}
	params.Attachments = []slack.Attachment{attachment}
	// params.AsUser = true
	params.Username = "연예봇"
	params.IconEmoji = ":+1:"
	channelID, timestamp, err := api.PostMessage(def.ChannelID, "이것을 원하셨습니까?", params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
