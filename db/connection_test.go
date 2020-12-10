package db

import (
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	c, err := ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", c)
}
