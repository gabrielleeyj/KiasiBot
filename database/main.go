package main

import "errors"

// User is the model representation of a user. Users are stored in the User Service.
type User struct {
	UUID string `json:"-"`
	Location
	Time
}

// Location is the representation of the users location.
type Location struct {
	Longitude, Latitude string `json:"longitude,latitude"`
}

// Time is the representation of the users time.
type Time struct {
	Day, Month, Year int `json:"day,month,year"`
	Hour, Min        int `json:"hour,min"`
}

// DataRepository is the interface that the data repository should conform to.
type DataRepository interface {
	Create(user User) (*User, error)
	AddLocation(ul Location) (*Location, error)
}

// repositoryMemory is used to store the data for users
type repositoryMemory struct {
	users User
}

func (r *repositoryMemory) Create(user User) (*User, error) {
	u := user
	if u != nil {
		return &u, nil
	}
	return nil, errors.New("user not created")
}

func (r *repositoryMemory) AddLocation(l Location) (*Location, error) {
	ul := l
	if ul != nil {
		return nil, errors.New("location not added")
	}
	return &ul, nil
}

func InMemoryRepository(user User) Repository {
	return &repositoryMemory{
		users: user,
	}
}
