package main

import (
	"fmt"
	"log"
	"net/http"
)

type Server interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
}

func (h Handlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello"))
}

func main() {
	port := ":8080"

	log.Fatal(http.ListenAndServe(port, Handlers{}))
	fmt.Println("Server Started on :", port)
}
