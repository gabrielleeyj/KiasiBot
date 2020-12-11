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

type ConnectionRepository interface {
	Session() (*mongo.Client, error)
	Database(name string) (*mongo.Database, error)
	Collection(name string) (*mongo.Collection, error)
}
type MongoDB struct {
	Client *mongo.Client
	DB     *mongo.Database
	Coll   *mongo.Collection
}

type Connection struct {
	app *MongoDB
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func Connect(database, collection string) (*mongo.Collection, error) {
	// initialize the client connection.
	conn := NewSession()
	c, err := conn.Session()
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

func NewSession() *MongoDB {
	return &MongoDB{}
}

// ConnectDB : helper function to connect to mongoDB database
func (m *MongoDB) Session() (*mongo.Client, error) {
	// main code to start connection
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_URI")))
	if err != nil {
		return nil, err
	}

	// ping cluster to check connection status
	err = client.Ping(context.TODO(), readpref.Primary())
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

	// No errors show success message
	fmt.Println("Connected to MongoDB Client!")

	// Connect to MongoDB collection for data storage
	// database := client.Database("db")
	return client, nil
}

func NewDatabase(client *mongo.Client) ConnectionRepository {
	return &MongoDB{Client: client}
}
func (m *MongoDB) Database(name string) (*mongo.Database, error) {
	db := m.Client.Database(name)
	fmt.Println("Connected to Datbase: ", name)
	return db, nil
}

func NewCollection(db *mongo.Database) ConnectionRepository {
	return &MongoDB{DB: db}
}

func (m *MongoDB) Collection(name string) (*mongo.Collection, error) {
	db := m.DB
	collection := db.Collection(name)
	fmt.Println("Connected to Collection: ", name)
	return collection, nil
}
