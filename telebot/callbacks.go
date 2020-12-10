package telebot

import tb "gopkg.in/tucnak/telebot.v2"

// makeButtons calls the function within tbot to create buttons for Telegram Chat
func makeButtons() *tb.ReplyMarkup {

	// shareLoc - Button to share location information
	shareLoc := tb.ReplyButton{
		Text:     "Share Location?",
		Location: true,
	}

	// getMap - returns URL of the Populated Map Data via Handler.
	getMap := tb.ReplyButton{
		Text: "Get Map",
	}

	return &tb.ReplyMarkup{
		ReplyKeyboard: [][]tb.ReplyButton{
			[]tb.ReplyButton{shareLoc},
			[]tb.ReplyButton{getMap}},
		ResizeReplyKeyboard: true,
		OneTimeKeyboard:     true,
	}
}
