package model

import (
	"KiasiBot/db"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/* this is the data model example for a get location
"result": [
                {
                        "update_id": 524107517, // text message id for editMessage to work
                        "message": {
                                "message_id": 205,
                                "from": {
                                        "id": <<int>>,
                                        "is_bot": false,
                                        "first_name": <<string>>,
                                        "username": <<string>>,
                                        "language_code": "en"
                                },
                                "chat": {
                                        "id": <<int>>,
                                        "first_name": <<string>>,
                                        "username": <<string>>,
                                        "type": "private"
                                },
                                "date": 1607481539, // stored in int
                                "location": {
                                        "latitude": 1.327737,
                                        "longitude": 103.889892
                                }
                        }
                }
		]
*/

// PostRepository sets the methods for manipulation of data
type PostRepository interface {
	GetAll() ([]Post, error)
	Create(post Post) (*Post, error)
}

// Post struct represents the structure of the data to Post.
type Post struct {
	ChatID    int64     `json:"-" bson:"ChatID,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt,omitempty" bson:"expiresAt"`
	Locations Location  `json:"locations,omitempty" bson:"locations"`
	Status    string    `json:"status,omitempty" bson:"status,omitempty"`
}

// Location is the data associated with the Post struct
type Location struct {
	Lat  float64 `json:"lat" bson:"lat"`
	Lng  float64 `json:"lng" bson:"lng"`
	Name string  `json:"name" bson:"name,omitempty"`
}

// PostStorage struct represents the collection and structure for the PostRepository.
type PostStorage struct {
	post Post
	col  *mongo.Collection
}

var (
	ctx        = context.TODO()
	database   = "db"
	collection = "usr"
)

// NewCreatePostRepository initializes the Create function to post into the database
func NewCreatePostRepository(post Post) PostRepository {
	return &PostStorage{post: post}
}

// Create method to post to database cloud
func (p *PostStorage) Create(post Post) (*Post, error) {
	// initialize the database connection.
	collection, err := db.Connect(database, collection)
	if err != nil {
		return nil, err
	}

	post.CreatedAt = time.Now()                          // logs time of creation
	post.ExpiresAt = time.Now().Add(time.Hour * 24 * 15) // adds 15 days from creation
	// Insert post to mongodb
	insertResult, err := collection.InsertOne(ctx, &post)
	if err != nil {
		return nil, err
	}
	fmt.Println("Inserted Post: ", insertResult.InsertedID)
	return &post, nil
}

// NewGetAllPostRepository initalizes the GetAll function to retrieve all the posts from the database
func NewGetAllPostRepository() PostRepository {
	return &PostStorage{}
}

// GetAll is a method to retrieve all documents in MongoDB and populate the data back into the memory.
func (p *PostStorage) GetAll() ([]Post, error) {

	// initialize the client connection.
	conn := db.NewSession()
	c, err := conn.Start()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	// initalize the database
	newdb := db.NewDatabase(c)
	d, err := newdb.Database("db")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(d)

	// initalize the collection
	newcollection := db.NewCollection(d)
	coll, err := newcollection.Collection("usr")
	if err != nil {
		log.Fatal(err)
	}

	// initialize the database connection.
	// collection, err := db.Connect(database, collection)
	// if err != nil {
	// 	return nil, err
	// }

	// bson.M{}, pass an empty filter to get all the data.
	cur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR: ", err)
		return nil, err
	}

	// defer after execution of a function until the surrounding function returns.
	// runs cur.Close() process after cur.Next().
	defer cur.Close(ctx)

	// initialize the post array
	posts := make([]Post, 0)

	// iterate through the cursor and deocode each entry
	for cur.Next(ctx) {

		// initializer to store the data
		var result Post

		// decodes the bson.D and maps it to the initializer
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		posts = append(posts, result)
	}
	// for testing
	fmt.Println(posts)

	return posts, nil
}
