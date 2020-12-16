package telebot

import (
	"KiasiBot/model"
	"fmt"
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

// StartBot invokes the connection command
func StartBot() {
	// bot settings
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL: "",
		// listen to the request / debugging purposes
		Verbose: true,
		Token:   os.Getenv("TELEGRAM_TOKEN"),
		Poller:  &tb.LongPoller{Timeout: 60 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Hello!", makeButtons())
	})

	b.Handle("Get Map", func(m *tb.Message) {
		b.Send(m.Sender, "You may view the mapdata 📍<a href=\"https://cutt.ly/WhJlc10\">here</a>", &tb.SendOptions{ParseMode: "HTML"})
	})

	// On reply button pressed (message)
	b.Handle(tb.OnLocation, func(m *tb.Message) {
		// if not private stop
		if !m.Private() {
			return
		}

		// casting of value to ensure compatibility with the db package
		lat := m.Location.Lat
		valuelat := float64(lat)

		lng := m.Location.Lng
		valuelng := float64(lng)

		// initialize the data model for posting
		data := model.Post{
			ChatID: m.Chat.ID,
			Locations: model.Location{
				Lat:  valuelat,
				Lng:  valuelng,
				Name: m.Chat.Username,
			},
			Status: "User",
		}

		fmt.Println(data)
		// initialize the PostRepository
		c := model.NewCreatePostRepository()
		c.Create(data)

		// return confirmation message
		b.Send(m.Sender, "Received Location")
		b.Send(m.Sender, "You may view the mapdata 📍<a href=\"https://cutt.ly/covid-chart\">here</a>", &tb.SendOptions{ParseMode: "HTML"})
	})

	// Starts the bot connection
	b.Start()
}
