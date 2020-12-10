package model

import (
	"testing"
)

func TestPost(t *testing.T) {

	// initialize post data
	newPost := Post{
		ChatID: 0,
		Locations: Location{
			Lat: 1.323850,
			Lng: 103.844560,
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

}
