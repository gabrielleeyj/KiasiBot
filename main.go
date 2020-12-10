package main

import (
	"KiasiBot/db"
	"KiasiBot/model"
	"KiasiBot/telebot"
	"fmt"
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

	// initialize the model to store posts
	posts := make([]model.Post, 0)
	dbM := model.NewGetHandler(posts)

	// call the function to Get Posts from MongoDB and stores to memeory
	dbM.GetPosts()

	fmt.Println(posts)
}
