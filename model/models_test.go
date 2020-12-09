package model

import (
	"testing"
	"time"
)

func TestPost(t *testing.T) {

	// initialize post data
	newPost := Post{
		ChatID:    123151231,
		CreatedAt: time.Now(),
		Locations: Location{
			Lat: 1.1111111,
			Lng: 2.2222222,
		},
		Status: "Test",
	}

	CreatePost(newPost)
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
