package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/krognol/go-wolfram"
	"github.com/shomali11/slacker"
	"github.com/tidwall/gjson"
	witai "github.com/wit-ai/wit-go/v2"
	"log"
	"os"
)

var wolframClient *wolfram.Client

func printCommandEvents(analytics <-chan *slacker.CommandEvent) {
	for events := range analytics {
		fmt.Println("command event")
		fmt.Println(events.Timestamp)
		fmt.Println(events.Parameters)
		fmt.Println(events.Command)
		fmt.Println(events.Event)
		fmt.Println()
	}

}
func main() {
	godotenv.Load(".env")
	bot := slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"), os.Getenv("SLACK_APP_TOKEN"))
	client := witai.NewClient(os.Getenv("WIT_AI_TOKEN"))
	wolframClient := &wolfram.Client{AppID: os.Getenv("WOLFRAM_APP_ID")}
	go printCommandEvents(bot.CommandEvents())

	bot.Command("<message>", &slacker.CommandDefinition{
		Description:       "send any question to wolfram",
		Example:           "What is color Red",
		AuthorizationFunc: nil,
		Handler: func(botCtx slacker.BotContext, request slacker.Request, response slacker.ResponseWriter) {
			query := request.Param("message")

			msg, _ := client.Parse(&witai.MessageRequest{
				Query: query,
			})
			data, _ := json.MarshalIndent(msg, "", "	")
			rough := string(data[:])
			value := gjson.Get(rough, "entities.wit$wolfram_search_query:wolfram_search_query.1.value").String()
			ans, err := wolframClient.GetSpokentAnswerQuery(value, wolfram.Metric, 100)
			if err != nil {
				ans = "Sorry can't response due to an error"
				log.Fatal(err)
			}
			response.Reply(ans)
		},
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
