package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yanzay/tbot"
)

type application struct {
	client *tbot.Client
}

var (
	app   application
	bot   *tbot.Server
	token string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	bot = tbot.New(token, tbot.WithLogger(tbot.BasicLogger{}))
	app.client = bot.Client()

	// // test response
	// bot.HandleMessage("", func(m *tbot.Message) {
	// 	app.client.SendMessage(m.Chat.ID, "hello!")
	// })
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/location", app.LocationHandler)
	bot.HandleMessage("/share", app.LocTestHandler)
	bot.HandleCallback(app.callBackHandler)
	log.Fatal(bot.Start())

}
