package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Location struct {
	// Latitude
	Lat float32 `json:"latitude"`
	// Longitude
	Lng float32 `json:"longitude"`

	// Period in seconds for which the location will be updated
	// (see Live Locations, should be between 60 and 86400.)
	LivePeriod int `json:"live_period,omitempty"`
}
type ReplyButton struct {
	Text string `json:"text"`

	Contact  bool `json:"request_contact,omitempty"`
	Location bool `json:"request_location,omitempty"`
}

func main() {
	// start bot
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		URL:    "",
		Token:  os.Getenv("TELEGRAM_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	menu := &ReplyButton{
		Contact:  false,
		Location: true,
	}

	b.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Hello!", menu)
	})

	// On reply button pressed (message)
	b.Handle(tb.OnLocation, func(m *tb.Message) {
		if !m.Private() {
			return
		}
		b.Send(m.Sender, "Received Location")
	})

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all the text messages that weren't
		// captured by existing handlers
	})

	b.Handle(tb.OnCallback, func(m *tb.Callback) {
		// incoming location messages
		// Location := Location{
		// 	Lat: m.Location.Lat,
		// 	Lng: m.Location.Lng,
		// }
		// fmt.Println(Location)
		b.Send(m.Sender, "location received")
	})

	b.Start()
}
