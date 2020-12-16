package server

import (
	"KiasiBot/model"
	"encoding/json"
	"html/template"
	"io/ioutil"
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

// FeatureCollection represents the collection of GeoJSON data
type FeatureCollection struct {
	Type     string    `json:"type"`
	Features []GeoJSON `json:"features"`
}

// GeoJSON data structure
type GeoJSON struct {
	Type string     `json:"type"`
	Geo  Geometry   `json:"geometry"`
	Prop Properties `json:"properties"`
}

// Geometry represents part of the GeoJSON structure
type Geometry struct {
	Type  string    `json:"type"`
	Coord []float64 `json:"coordinates"`
}

// Properties represents part of the GeoJSON structure
type Properties struct {
	Name string `json:"name"`
}

// Map represents the data structure to store the resp from GetAll()
type Map struct {
	Data []GeoJSON
}

// Presenter represents the functions that can be executed
type Presenter interface {
	Home() http.HandlerFunc
	CSS(dir string) http.HandlerFunc
	JavaScript(dir string) http.HandlerFunc
	JSON(dir string) http.HandlerFunc
}

func parseTemplate(w http.ResponseWriter, file string) error {
	t, err := template.ParseGlob(file)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return err
	}
	// Execute Function on Load
	GetData()

	err = t.Execute(w, t)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return err
	}
	return nil
}

type presenter struct{}

func (p *presenter) Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		err := parseTemplate(w, "index.gohtml")
		if err != nil {
			return
		}
	}
}
func (p *presenter) JSON(dir string) http.HandlerFunc {
	return http.StripPrefix("/json", http.FileServer(http.Dir(dir))).ServeHTTP
}

func (p *presenter) CSS(dir string) http.HandlerFunc {
	return http.StripPrefix("/css", http.FileServer(http.Dir(dir))).ServeHTTP
}

func (p *presenter) JavaScript(dir string) http.HandlerFunc {
	return http.StripPrefix("/js", http.FileServer(http.Dir(dir))).ServeHTTP
}

// NewPresenter represents part of the Interface implementation
func NewPresenter() Presenter {
	return &presenter{}
}

// GetData handles the request for getting Loc Data from MongoDB
func GetData() error {
	// initialize GetAll
	mongodb := model.NewGetAllPostRepository()

	// call the function to Get Posts from MongoDB
	data, err := mongodb.GetAll()
	if err != nil {
		return err
	}

	MapData := make([]GeoJSON, 0)

	for _, v := range data {
		GeoJSON := GeoJSON{
			Type: "Feature",
			Geo: Geometry{
				Type: "Point",
				Coord: []float64{
					v.Locations.Lng,
					v.Locations.Lat,
				},
			},
			Prop: Properties{
				Name: v.Status,
			},
		}
		MapData = append(MapData, GeoJSON)
	}
	FeatureCollection := FeatureCollection{
		Type:     "FeatureCollection",
		Features: MapData,
	}
	file, err := json.MarshalIndent(FeatureCollection, "", " ")
	if err != nil {
		return err
	}
	_ = ioutil.WriteFile("./static/json/pointer.geojson", file, 0644)
	return nil
}
