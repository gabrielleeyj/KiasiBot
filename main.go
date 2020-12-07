package main

import (
	"KiasiBot/db"
	"KiasiBot/model"
	"fmt"
	"time"
)

func main() {
	db.ConnectDB()
	newPost := model.Post{
		ChatID:    "test",
		CreatedAt: time.Now(),
		Locations: model.Location{
			Lat: 1.1111111,
			Lng: 2.2222222,
		},
	}
	fmt.Println(newPost)
	fmt.Println(model.CreatePost(newPost))

}
