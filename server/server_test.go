package main

import (
	"fmt"
	"log"
	"testing"
)

func TestGetData(t *testing.T) {
	data, err := GetData()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data)
}
