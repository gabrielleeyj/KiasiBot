package model

import (
	"fmt"
	"log"
	"testing"
)

// func TestCreate(t *testing.T) {

// 	// initialize post data
// 	newPost := Post{
// 		ChatID: 0,
// 		Locations: Location{
// 			Lat: 1.323850,
// 			Lng: 103.844560,
// 		},
// 		Status: "Test",
// 	}

// 	// initialize NewCreatePostRepository with params
// 	mongo := NewCreatePostRepository(newPost)

// 	// call the function
// 	mongo.Create(newPost)

// }

func TestGetAll(t *testing.T) {
	// initialize GetAll
	mongodb := NewGetAllPostRepository()

	posts, err := mongodb.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(posts)
	// call the function to Get Posts from MongoDB

}
