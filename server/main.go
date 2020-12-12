package main

import (
	"KiasiBot/model"
	"fmt"
	"log"
	"net/http"
)

func main() {

	// handle server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initializedata() {
	// initialize GetAll
	mongodb := model.NewGetAllPostRepository()

	// call the function to Get Posts from MongoDB
	data, err := mongodb.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range data {
		fmt.Println("Locations: ", v.Locations)
		fmt.Println("Status: ", v.Status)
	}
}
