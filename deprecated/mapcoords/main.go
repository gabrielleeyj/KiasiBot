package mapcoords

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"log"
)

// API key
const client := geo.NewGoogleGeo("api-key")

// This struct contains selected fields from Google's Geocoding Service response
type geolocateResponse struct {
	Error struct {
		Code    int
		Message string
		Errors  []struct {
			Domain  string
			Reason  string
			Message string
		}
	}
	Location struct {
		Lat float64
		Lng float64
	}
}

// getLoc function returns the current location information of the client.
func getLoc() (*Location,error){
	res, err := client.Geolocate()
	if err != nil {
		return &res, nil
	}
	return nil, errors.New("cannot find location")
}


// Geolocate returns a rough location based on your IP
func (g *GoogleGeo) Geolocate() (*Point, error) {
	data, err := g.geolocateRequest()
	if err != nil {
		return nil, err
	}

	res := &geolocateResponse{}
	json.Unmarshal(data, res)
	if res.Error.Code != 0 {
		e := res.Error.Errors[0]
		return nil, fmt.Errorf(e.Domain + "." + e.Reason + "." + e.Message)
	}

	return &Point{Lat: res.Location.Lat, Lng: res.Location.Lng}, nil
}

func (g *GoogleGeo) geolocateRequest() ([]byte, error) {
	if g.apiKey == "" {
		return nil, fmt.Errorf("Google API key not provided")
	}
	dst := "https://www.googleapis.com/geolocation/v1/geolocate?key=" + g.apiKey
	form := url.Values{}

	req, err := http.NewRequest("POST", dst, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, err
	}
	resp, requestErr := g.client.Do(req)
	if requestErr != nil {
		return nil, requestErr
	}
	data, dataReadErr := ioutil.ReadAll(resp.Body)
	if dataReadErr != nil {
		return nil, dataReadErr
	}
	return data, nil
}