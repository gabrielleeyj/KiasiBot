package main

import (
	"KiasiBot/model"
	"html/template"
	"log"
	"net/http"
)

/* Example DataSet for a GeoJSON data
{
  "type": "Feature",
  "geometry": {
    "type": "Point",
    "coordinates": [125.6, 10.1]
  },
  "properties": {
    "name": "Dinagat Islands"
  }
}
*/

// Map represents the data structure to store the resp from GetAll()
type Map struct {
	Data []model.Post
}

func main() {

	// handle server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		t, err := template.ParseGlob("*.gohtml")
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		// on request GetData() should retrieve the data from MongoDB and store it to the memory
		data, err := GetData()
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})

	// start http server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// GetData handles the request for getting Loc Data from MongoDB
func GetData() (*Map, error) {
	// initialize GetAll
	mongodb := model.NewGetAllPostRepository()

	// call the function to Get Posts from MongoDB
	data, err := mongodb.GetAll()
	if err != nil {
		return nil, err
	}
	// for _, v := range data {
	// 	fmt.Println("Locations: ", v.Locations)
	// 	fmt.Println("Status: ", v.Status)
	// }

	return &Map{Data: data}, nil
}
