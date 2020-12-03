package main

import (
	"fmt"
	"testing"
)

func testCreate(t *testing.T) {
	u := User{
		UUID:     "test",
		Location: Location{Longitude: "234", Latitude: "123"},
		Time:     Time{Day: 1, Month: 12, Year: 2020, Hour: 23, Min: 59},
	}
	result := Create(u)
	if result != nil {
		fmt.Printf("Result: %v", result)
	}
}
