package model

import (
	"testing"
	"time"
)

func TestPost(t *testing.T) {

	// initialize post data
	newPost := Post{
		ChatID:    "test",
		CreatedAt: time.Now(),
		Locations: Location{
			Lat: 1.1111111,
			Lng: 2.2222222,
		},
	}

	CreatePost(newPost)
}
