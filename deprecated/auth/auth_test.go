package auth

import (
	"fmt"
	"testing"
)

func TestEncode(t *testing.T) {
	body := Data{
		ChatID: "1",
		Locations: Location{
			Lat: 1.312451241312,
			Lng: 103.876576657,
		},
	}

	result, time, err := Encode(body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Result %v\n", result)
	fmt.Printf("Time %v\n", time)

}
