package main

import (
	"github.com/yanzay/tbot/v2"
)

func (a *application) startHandler(m *tbot.Message) {
	msg := "This is a bot whose sole purpose is to play, rock, paper, scissors with you. Use /play to play."
	a.client.SendMessage(m.Chat.ID, msg)
}

func (a *application) playHandler(m *tbot.Message) {
	buttons := makeButtons()
	a.client.SendMessage(m.Chat.ID, "Pick an option:", tbot.OptInlineKeyboardMarkup(buttons))
}

func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	humanMove := cq.Data
	msg := draw(humanMove)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}

