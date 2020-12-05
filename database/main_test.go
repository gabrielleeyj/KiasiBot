package db

import (
	"fmt"
	"log"
	"testing"
)

func TestCreateUser(t *testing.T) {
	users := make([]User, 0)
	dbM := NewRepository(users)

	newUser := User{
		ID: 1,
		Locations: Location{
			Latitude:  1.111111,
			Longitude: 1.111111,
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
		ID: 1,
		Locations: Location{
			Latitude:  1.111111,
			Longitude: 1.111111,
		},
	}
	c, _ := dbM.Create(newUser)
	if c != nil {
		fmt.Println("User Created - ", c)
	}

	r := dbM.Delete(1)
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
		ID: 1,
		Locations: Location{
			Latitude:  1.111111,
			Longitude: 1.111111,
		},
	}
	c, _ := dbM.Create(newUser)
	if c != nil {
		fmt.Println("User Created - ", c)
	}

	newUser2 := User{
		ID: 1,
		Locations: Location{
			Latitude:  2.222222,
			Longitude: 2.2222222,
		},
	}
	r, _ := dbM.Update(1, newUser2)
	if r != nil {
		fmt.Println("User Updated - ", r)
	}

}
