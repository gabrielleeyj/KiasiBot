package coords

// Coords struct stores the coordinates for googlemaps api
type Coords struct {
	lat float64
	lng float64
}

// data represented in the past 14 days from
// https://www.gov.sg/article/covid-19-public-places-visited-by-cases-in-the-community-during-infectious-period
// populated via their coords declared in the const var : struct coords
var (
	SQUE = Coords{
		lat: 1.289201,
		lng: 103.846263,
	}
	MAKANSUTRAGB = Coords{
		lat: 1.289393,
		lng: 103.856870,
	}
	MBSCASINO = Coords{
		lat: 1.283883,
		lng: 103.860744,
	}
	CAKOPITIAM = Coords{
		lat: 1.355446,
		lng: 103.986313,
	}
	S11BISHAN = Coords{
		lat: 1.349235,
		lng: 103.848417,
	}
	WILKE = Coords{
		lat: 1.302094,
		lng: 103.848829,
	}
	DOCTP = Coords{
		lat: 1.279642,
		lng: 103.844026,
	}
)
