package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectionRepository manages all the methods required for a connection
type ConnectionRepository interface {
	Start() (*mongo.Client, error)
	Database(name string) (*mongo.Database, error)
	Collection(name string) (*mongo.Collection, error)
}

// MongoDB implements the ConnectionRepository interface.
type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
}

// init checks for the .env file
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

// Connect combines the layers required to connect MongoDB
func Connect(database, collection string) (*mongo.Collection, error) {
	// initialize the client connection.
	conn := NewSession()
	c, err := conn.Start()
	if err != nil {
		return nil, err
	}

	// initalize the database
	newdb := NewDatabase(c)
	db, err := newdb.Database(database)
	if err != nil {
		return nil, err
	}

	// initalize the collection
	newcollection := NewCollection(db)
	coll, err := newcollection.Collection(collection)
	if err != nil {
		return nil, err
	}
	return coll, nil
}

// NewSession starts the mongoDB client session
func NewSession() *MongoDB {
	return &MongoDB{}
}

// Start connects to mongoDB database client layer
func (m *MongoDB) Start() (*mongo.Client, error) {
	// main code to start connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	// ping cluster to check connection status
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	fmt.Println("Ping:", err)

	// check available databases and prints output
	// databases, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(databases)
	return client, nil
}

// NewDatabase starts a new database connection
func NewDatabase(client *mongo.Client) ConnectionRepository {
	return &MongoDB{Client: client}
}

// Database returns the database connection
func (m *MongoDB) Database(name string) (*mongo.Database, error) {
	db := m.Client.Database(name)
	return db, nil
}

// NewCollection starts a new collection
func NewCollection(db *mongo.Database) ConnectionRepository {
	return &MongoDB{DB: db}
}

// Collection returns the collection connection
func (m *MongoDB) Collection(name string) (*mongo.Collection, error) {
	db := m.DB
	collection := db.Collection(name)
	return collection, nil
}
