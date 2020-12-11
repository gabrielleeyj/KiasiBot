package db

import (
	"fmt"
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	connection, err := Connect("db", "usr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(connection)
}

func TestSession(t *testing.T) {
	// initialize the client connection.
	conn := NewSession()
	c, err := conn.Session()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func TestDatabase(t *testing.T) {
	// initialize the client connection.
	conn := NewSession()
	c, err := conn.Session()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	// initalize the database
	newdb := NewDatabase(c)
	db, err := newdb.Database("db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)
}

func TestCollection(t *testing.T) {
	// initialize the client connection.
	conn := NewSession()
	c, err := conn.Session()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	// initalize the database
	newdb := NewDatabase(c)
	db, err := newdb.Database("db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	// initalize the collection
	newcollection := NewCollection(db)
	coll, err := newcollection.Collection("usr")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(coll)
}
