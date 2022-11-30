package main

import (
	"log"

	"github.com/slack-go/slack"
)

var (
	SLACK_BOT_TOKEN = "xoxb-4014835228961-4447289789300-K6WJEg94Ri0E7vPMTNGY2yQO"
	CHANNEL_ID = "C040231LT0B" // Postric Random Channel
)

func main()  {
	log.Println("Launching bot")

	api := slack.New(SLACK_BOT_TOKEN, slack.OptionDebug(true))
	channels := []string{CHANNEL_ID}
	files := []string{}

	for _, file := range files {
		params := slack.FileUploadParameters{
			Channels: channels,
			File: file,
		}

		f, err := api.UploadFile(params)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Printf("file name: %s and file url: %s", f.Name, f.URLPrivate)
	}
}
