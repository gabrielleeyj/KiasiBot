package main

import (
	"KiasiBot/db"
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
	db.ConnectDB()
	telebot.StartBot()

	// newPost := model.Post{
	// 	ChatID:    "test",
	// 	CreatedAt: time.Now(),
	// 	Locations: model.Location{
	// 		Lat: 1.1111111,
	// 		Lng: 2.2222222,
	// 	},
	// }
	// fmt.Println(newPost)
	// fmt.Println(model.CreatePost(newPost))

}
