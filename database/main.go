package db

import (
	"errors"
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
