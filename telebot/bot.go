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
		URL:     "",
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
		b.Send(m.Sender, "You may view the mapdata 📍<a href=\"https://cutt.ly/covid-chart\">here</a>", &tb.SendOptions{ParseMode: "HTML"})
	})

	// On reply button pressed (message)
	b.Handle(tb.OnLocation, func(m *tb.Message) {
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
		// c := model.NewCreatePostRepository(context.TODO(), data)
		// c.Create(context.TODO(), data)
		// return confirmation message
		b.Send(m.Sender, "Received Location")
		b.Send(m.Sender, "You may view the mapdata 📍<a href=\"https://cutt.ly/covid-chart\">here</a>", &tb.SendOptions{ParseMode: "HTML"})
	})

	// Starts the bot connection
	b.Start()
}
