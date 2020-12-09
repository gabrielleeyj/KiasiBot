package main

import (
	"KiasiBot/model"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/tucnak/telebot.v2"
)

// initialize check for .env file existence
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func main() {
	// bot settings
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL:     "",
		Verbose: true,
		Token:   os.Getenv("TELEGRAM_TOKEN"),
		Poller:  &tb.LongPoller{Timeout: 60 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	go b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Hello!", makeButtons())
	})

	go b.Handle("Get Map", func(m *tb.Message) {
		b.Send(m.Sender, "You may view the mapdata 📍<a href=\"https://cutt.ly/covid-chart\">here</a>", &tb.SendOptions{ParseMode: "HTML"})
	})

	// On reply button pressed (message)
	go b.Handle(tb.OnLocation, func(m *tb.Message) {
		// if not private stop
		if !m.Private() {
			return
		}
		data := model.Post{
			ChatID: m.Chat.ID,
			Locations: model.Location{
				Lat:  m.Location.Lat,
				Lng:  m.Location.Lng,
				Name: m.Chat.Username,
			},
			Status: "User",
		}

		fmt.Println(data)
		model.CreatePost(data)
		// return confirmation message
		b.Send(m.Sender, "Received Location")
		b.Send(m.Sender, "You may view the mapdata 📍<a href=\"https://cutt.ly/covid-chart\">here</a>", &tb.SendOptions{ParseMode: "HTML"})
	})

	// Starts the bot connection
	b.Start()
}
