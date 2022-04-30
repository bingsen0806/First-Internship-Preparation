package main

import (
	"fmt"
	"os"

	"github.com/bingsen0806/slack-file-bot/secret"
	"github.com/slack-go/slack"
)

func main() {
	os.Setenv("SLACK_BOT_TOKEN", secret.BotToken)
	os.Setenv("CHANNEL_ID", secret.ChannelId)
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"slackfile.txt"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters {
			Channels: channelArr,
			File: fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name: %s, URL:%s\n", file.Name, file.URL)
	}
}