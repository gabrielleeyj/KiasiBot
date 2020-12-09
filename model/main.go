package model

import (
	"KiasiBot/db"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

/* this is the data model example for get location
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

// UserRepository sets the methods for manipulation of data
type UserRepository interface {
	GetPosts() error
}

// Post struct represents the structure of the data to Post.
type Post struct {
	ChatID    int64     `json:"-" bson:"ChatID,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	ExpiresAt time.Time `json:"expiresAt,omitempty" bson:"expiresAt"`
	Locations Location
	Status    string `json:"status,omitempty" bson:"status,omitempty"`
}

// Location is the data associated with the Post struct
type Location struct {
	Lat  float32 `json:"lat" bson:"lat"`
	Lng  float32 `json:"lng" bson:"lng"`
	Name string  `json:"name" bson:"name,omitempty"`
}

type dbMemory struct {
	posts []Post
}

// CreatePost method to post to database cloud
func CreatePost(post Post) error {

	// Connect to Collection connection
	c := db.ConnectDB()

	// set default mongodb ID  and created date

	post.CreatedAt = time.Now()                          // logs time of creation
	post.ExpiresAt = time.Now().Add(time.Hour * 24 * 15) // adds 15 days from creation
	// Insert post to mongodb
	insertResult, err := c.InsertOne(context.TODO(), post)
	if err != nil {
		return err
	}
	fmt.Println("Inserted Post: ", insertResult.InsertedID)
	return nil
}

func NewGetHandler(initial []Post) UserRepository {
	return &dbMemory{posts: initial}
}

func (d *dbMemory) GetPosts() error {
	// Connect to Collection connection
	c := db.ConnectDB()

	// bson.D{}, pass empty filter to get all the data.
	cur, err := c.Find(context.TODO(), bson.D{})
	if err != nil {
		fmt.Println("Finding all documents ERROR: ", err)
	}

	// defer after execution of a function until the surrounding function returns.
	// runs cur.Close() process after cur.Next().
	defer cur.Close(context.TODO())

	// iterate through the cursor and deocode each entry
	for cur.Next(context.TODO()) {
		var result bson.M

		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

	}

	return nil
}
