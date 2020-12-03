package main


var picks = []string{"rock","paper","scissors"},

func init() {
	rand.Seed(time.Now().Unix())
}

func makeButtons() *tbot.InlineKeyboardMarkup {
	btnRock := tbot.InlineKeyboardButton {
		Text: "Rock",
		CallbackData: "rock",
	}

	btnPaper := tbot.InlineKeyboardButton {
		Text: "Paper",
		CallbackData: "paper",
	}

	btnScissors := tbot.InlineKeyboardButton {
		Text: "Scissors",
		CallbackData: "scissors",
	}

	return &tbot.InlineKeyboardMarkup {
		InlineKeyboard: [][]tbot.InlineKeyboardButton {
			[]tbot.InlineKeyboardButton{btnRock, btnPaper, btnScissors},
		}
	}
}

func draw(humanMove string) (msg string) {
	var result string
	botMove := picks[rand.Intn(len(picks))]

	// Determine outcome

switch humanMove {
case botMove:
	result = "draw"
case options[botMove]:
	result = "lost"
default:
	result = "won"
}
msg = fmt.Sprintf("You %s! You chose %s and I chose %s", result, humanMove, botMove)
return
}