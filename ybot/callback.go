package main

import "github.com/yanzay/tbot"

var options = []string{"Info", "GetLocation", "ShareLocation"}
var info = "This bot is built with Golang"

func makeButtons() *tbot.InlineKeyboardMarkup {
	btnInfo := tbot.InlineKeyboardButton{
		Text:         "Info",
		CallbackData: info,
	}
	btnGetLoc := tbot.InlineKeyboardButton{
		Text:         "Get",
		CallbackData: "Shows Map Data",
	}
	btnShareLoc := tbot.InlineKeyboardButton{
		Text:         "Share",
		CallbackData: "Shares Location & returns Map Data",
	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnInfo, btnGetLoc, btnShareLoc},
		},
	}
}

func locButton() *tbot.ReplyKeyboardMarkup {
	btnLocation := tbot.KeyboardButton{Text: "Give me your location!", RequestLocation: true}

	return &tbot.ReplyKeyboardMarkup{
		Keyboard: [][]tbot.KeyboardButton{
			[]tbot.KeyboardButton{btnLocation},
		},
		OneTimeKeyboard: true,
	}
}
