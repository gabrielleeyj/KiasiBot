package main

import tb "gopkg.in/tucnak/telebot.v2"

func makeButtons() *tb.ReplyMarkup {

	shareLoc := tb.ReplyButton{
		Text:     "Share Location?",
		Location: true,
	}

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
