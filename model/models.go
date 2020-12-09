package model

import (
	"KiasiBot/db"
	"context"
	"fmt"
	"time"
)

type Post struct {
	ChatID    int64     `json:"-" bson:"ChatID,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt,omitempty" bson:"expiresAt"`
	Locations Location
	Status    string `json:"status,omitempty" bson:"status,omitempty"`
}

type Location struct {
	Lat  float32 `json:"lat" bson:"lat"`
	Lng  float32 `json:"lng" bson:"lng"`
	Name string  `json:"name" bson:"name,omitempty"`
}

// func NewPost(p Post) *Post {
// 	return &Post{}
// }

// CreatePost method to post to databse
func CreatePost(post Post) error {

	// Get post collection connection
	c := db.ConnectDB()

	// set default mongodb ID  and created date

	post.CreatedAt = time.Now()                          // logs time of creation
	post.ExpiresAt = time.Now().Add(time.Hour * 24 * 15) // adds 15 days from creation
	// Insert post to mongodb
	insertResult, err := c.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	fmt.Println("Inserted Post: ", insertResult.InsertedID)
	return nil
}
