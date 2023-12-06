package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/shomali11/slacker"
)

func printCommandEvents(commandsEvents <-chan *slacker.CommandEvent) {
	for event := range commandsEvents {
		fmt.Println("Commads Events")
		fmt.Println(event.Timestamp)
		fmt.Println(event.Command)
		fmt.Println(event.Parameters)
		// fmt.Println(event.Event)
		fmt.Println()
	}
}

func ageCalculator(botCtx slacker.BotContext, req slacker.Request, res slacker.ResponseWriter) {
	currentYear := time.Now().Year()
	year := req.Param("year")
	yob, err := strconv.Atoi(year)
	if err != nil {
		res.Reply("Invalid Year provided")
		return
	}

	if yob >= currentYear {
		// r := fmt.Sprintf("Invalid Data")
		res.Reply("Invalid Year provided")
		return
	}
	age := currentYear - yob
	r := fmt.Sprintf("age is %d", age)
	res.Reply(r)

}

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Some error occured. Err: %s", err)
	// }

	APP_TOKEN := os.Getenv("SLACK_APP_TOKEN")
	BOT_TOKEN := os.Getenv("SLACK_BOT_TOKEN")
	//this func will create a new slack client
	bot := slacker.NewClient(BOT_TOKEN, APP_TOKEN)

	// a go routine to print commands....
	go printCommandEvents(bot.CommandEvents())

	// the logic for slack age calculator......
	bot.Command("My yob is <year>", &slacker.CommandDefinition{
		Description: "Age calculator",
		Examples:    []string{"My yob is 2020"},
		Handler:     ageCalculator,
	})

	// creating a context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
