package model

import (
	"KiasiBot/db"
	"context"
	"fmt"
	"time"
)

type Post struct {
	ChatID    string    `json:"-" bson:"ChatID,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at"`
	Locations Location
}

type Location struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

// func NewPost(p Post) *Post {
// 	return &Post{}
// }

// CreatePost method to post to databse
func CreatePost(post Post) error {

	// Get post collection connection
	c := db.ConnectDB()

	// set default mongodb ID  and created date

	post.CreatedAt = time.Now()

	// Insert post to mongodb
	insertResult, err := c.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	fmt.Println("Inserted Post: ", insertResult.InsertedID)
	return nil
}
