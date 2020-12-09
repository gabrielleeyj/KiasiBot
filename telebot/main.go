package main

import (
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
		Poller:  &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	menu := tb.ReplyButton{
		Text:     "Share Location?",
		Location: true,
	}

	getMap := tb.ReplyButton{
		Text: "Get Map",
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Hello!", &tb.ReplyMarkup{
			ReplyKeyboard:       [][]tb.ReplyButton{[]tb.ReplyButton{menu}, []tb.ReplyButton{getMap}},
			ResizeReplyKeyboard: true,
			OneTimeKeyboard:     true,
		},
		)
	})

	b.Handle("Get Map", func(m *tb.Message) {
		b.Send(m.Sender, "https://cutt.ly/covid-chart")
	})
	// On reply button pressed (message)
	b.Handle(tb.OnLocation, func(m *tb.Message) {
		// if not private stop
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Received Location")
	})

	// Starts the bot connection
	b.Start()
}
