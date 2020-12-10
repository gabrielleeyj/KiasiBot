package model

import (
	"testing"
	"time"
)

func TestPost(t *testing.T) {

	// initialize post data
	newPost := Post{
		ChatID:    123151231,
		CreatedAt: time.Now(),
		Locations: Location{
			Lat: 1.1111111,
			Lng: 2.2222222,
		},
		Status: "Test",
	}

	CreatePost(newPost)
}

func TestGet(t *testing.T) {
	// initialize post model
	posts := make([]Post, 0)
	dbM := NewGetHandler(posts)

	// call the function to Get Posts from MongoDB
	dbM.GetPosts()
	// Test print
	// fmt.Println("Data: ", r)
	// return r
}
