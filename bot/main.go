package main

import (
	"log"
	"os"
)

type application struct {
	client *tbot.Client
}

var (
	app     application
	bot     *tbot.Server
	token   string
	options = map[string]string{
		"paper":    "rock",
		"rock":     "scissors",
		"scissors": "paper",
	}
)

func init() {
	e := godotenv.Load()
	if e != nil {
		log.Println(e)
	}
	token = os.Getenv("TELEGRAM_TOKEN")
}

func main() {
	bot = tbot.New(token, tbot.WithWebHook("https://", ":"+os.Getenv("PORT")))
	app.Client = bot.Client()
	bot.HandleMessage("/start", app.startHandler)
	bot.HandleMessage("play", app.playHandler)
	bot.HandleCallback(app.callbackHandler)
	log.Fatal(bot.Start())
}
