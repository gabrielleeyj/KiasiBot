package server

import (
	"log"
	"testing"
)

func TestGetData(t *testing.T) {
	err := GetData()
	if err != nil {
		log.Fatal(err)
	}

}
