package memory

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	users := make([]User, 0)
	dbM := NewRepository(users)

	newUser := User{
		ChatID: "1",
		Locations: Location{
			Lat: 1.111111,
			Lng: 1.111111,
		},
	}

	r, err := dbM.Create(newUser)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Test Passed - ", r)
}
func TestDeleteUser(t *testing.T) {
	users := make([]User, 0)
	dbM := NewRepository(users)
	fmt.Println("User Memory Repository Initialized - ", users)

	newUser := User{
		ChatID: "1",
		Locations: Location{
			Lat: 1.111111,
			Lng: 1.111111,
		},
	}
	c, _ := dbM.Create(newUser)
	if c != nil {
		fmt.Println("User Created - ", c)
	}

	r := dbM.Delete("1")
	if r != nil {
		fmt.Println("Test Passed - ", r)
	}
	fmt.Println("User Deleted - ", r)
}

func TestUpdateUser(t *testing.T) {
	users := make([]User, 0)
	dbM := NewRepository(users)
	fmt.Println("User Memory Repository Initialized - ", users)

	newUser := User{
		ChatID: "1",
		Locations: Location{
			Lat: 1.111111,
			Lng: 1.111111,
		},
	}
	c, _ := dbM.Create(newUser)
	if c != nil {
		fmt.Println("User Created - ", c)
	}

	newUser2 := User{
		ChatID: "2",
		Locations: Location{
			Lat: 2.222222,
			Lng: 2.2222222,
		},
	}
	r, _ := dbM.Update("1", newUser2)
	if r != nil {
		fmt.Println("User Updated - ", r)
	}

}
