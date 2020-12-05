package db

import (
	"errors"
	"fmt"
)

// UserRepository Interface represents the underlying implementations for the database driver
type UserRepository interface {
	Create(user User) (*User, error)
	Delete(id int) error
	Update(id int, user User) (*User, error)
}

// User is created when the Telegram Bot starts for a new user.
type User struct {
	ID        int `json:"-"`
	Locations Location
}

// Location stores the GPS coordinates from the user
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// dbMemory stores user data in memory
type dbMemory struct {
	users []User
}

func (d *dbMemory) Create(user User) (*User, error) {
	d.users = append(d.users, user)
	return &user, nil
}

func (d *dbMemory) Delete(id int) error {

	for idx, u := range d.users {
		if u.ID == id {
			d.users = append(d.users[:idx], d.users[idx+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}

func (d *dbMemory) Update(id int, user User) (*User, error) {
	for idx, u := range d.users {
		if u.ID == id {
			d.users[idx] = user
		}
	}
	return &user, nil
}

// NewRepository initializer which stores users data
func NewRepository(initial []User) UserRepository {
	return &dbMemory{
		users: initial,
	}
}

func main() {
	// initialize memory
	users := make([]User, 0)
	dbM := NewRepository(users)
	fmt.Println("Test [001] === Initialize", dbM)

	newUser := User{
		ID: 1,
		Locations: Location{
			Latitude:  1.111111,
			Longitude: 1.111111,
		},
	}
	dbM.Create(newUser)
	fmt.Println("Test [002] === Add User 1", dbM)
	newUser2 := User{
		ID: 2,
		Locations: Location{
			Latitude:  2.222222,
			Longitude: 2.222222,
		},
	}

	dbM.Create(newUser2)
	fmt.Println("Test [003] === Add User 2", dbM)

	dbM.Delete(1)
	fmt.Println("Test [004] === Removed 1 ", dbM)

	dbM.Create(newUser)
	fmt.Println("Test [005] === Add User 1", dbM)

	dbM.Update(1, newUser2)
	fmt.Println("Test [006] === Update User 1 to User 2 ", dbM)
}
