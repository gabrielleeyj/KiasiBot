package main

import (
	"log"
	"os"

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
	// e := godotenv.Load()
	// if e != nil {
	// 	log.Println(e)
	// }
	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	bot = tbot.New(token)
	app.client = bot.Client()

	// test response
	bot.HandleMessage("", func(m *tbot.Message) {
		app.client.SendMessage(m.Chat.ID, "hello!")
	})

	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("/location", app.LocationHandler)
	bot.HandleMessage("/share", app.LocTestHandler)
	// bot.HandleInlineResult(app.InlineQueryHandler)
	bot.HandleCallback(app.callBackHandler)
	log.Fatal(bot.Start())
}
