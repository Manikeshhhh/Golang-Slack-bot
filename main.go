package main

import(
	"fmt"
	"github.com/shomali11/slacker"
	"os"
	"context"
	"log"
)


func printCommandEvents(notes <-chan *slacker.CommandEvent){
	for event := range notes {
		fmt.Println("Command events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		fmt.Println(event.Event)
		fmt.Println()
	}

}

func main(){

	os.Setenv("SLACK_BOT_TOKEN","xoxb-********-*******-**********")
	os.Setenv("SLACK_APP_TOKEN","xapp-1-***********-********-*******************************")


	bot:= slacker.NewClient(os.Getenv("SLACK_BOT_TOKEN"),os.Getenv("SLACK_APP_TOKEN"))

go printCommandEvents(bot.CommandEvents())

	bot.Command("ping",&slacker.CommandDefinition{
		Handler : func(botctx slacker.BotContext,request slacker.Request,response slacker.ResponseWriter){
			response.Reply("pong")
		},
	})

	ctx,cancel:= context.WithCancel(context.Background())
	defer cancel()

	err:=bot.Listen(ctx)
	if err != nil {
		log.Fatal(err) 
	}
}