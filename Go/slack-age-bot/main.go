package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/bingsen0806/slack-age-bot/secret"
	"github.com/shomali11/slacker"
)

func printCommandEvents(analysticsChannel <- chan *slacker.CommandEvent) {
	for event := range analysticsChannel {
		fmt.Println("Command Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}
}
func main() {
	os.Setenv("SLACK_BOT_TOKEN", secret.BotToken)
	os.Setenv("SLACK_APP_TOKEN", secret.AppToken)

	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	go printCommandEvents(bot.CommandEvents())

	bot.Command("my yob is <year>", &slacker.CommandDefinition {
		Description: "yob calculator",
		Example: "my yob is 2020",
		Handler: func (botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			year := request.Param("year")
			yob, err := strconv.Atoi(year)
			if err != nil {
				print("error")
			}
			age := 2022-yob
			r := fmt.Sprintf("Age is %d", age)
			response.Reply(r)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}