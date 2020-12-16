package main

import (
	"KiasiBot/telebot"
	"log"

	"github.com/joho/godotenv"
)

// initialize check for .env file existence
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}
func main() {
	// start the bot server
	telebot.StartBot()

}
