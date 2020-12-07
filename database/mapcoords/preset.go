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

// EMBEDED
// https://www.onemap.sg/amm/amm.html?mapStyle=Default&zoomLevel=15&marker=latLng:1.2843237295033,103.85956355313499!iwt:JTNDcCUzRTI3JTIwTm92JTIwLSUyMDE0MjBoJTIwdG8lMjAxOTAwaCUyMC0lMjBNQlMlMjBDYXNpbm8lM0MlMkZwJTNF!icon:fa-child!colour:red&marker=latLng:1.3020840794266,103.849621716147!iwt:JTNDcCUzRTI1JTIwTm92JTIwLSUyMDE1MzBoJTIwdG8lMjAxODAwaCUyMC0lMjBXaWxraWUlMjBFZGdlJTNDJTJGcCUzRQ==!icon:fa-child!colour:red&marker=latLng:1.35671611225461,103.986514607903!iwt:JTNDcCUzRTI2JTIwTm92JTIwMjA1NWglMjB0byUyMDIxNTVoJTIwLSUyMEtvcGl0aWFtJTNDJTJGcCUzRSUwQSUzQ3AlM0UyNCUyME5vdiUyMDIxMDVoJTIwdG8lMjAyMjAwaCUyMC0lMjBLb3BpdGlhbSUzQyUyRnAlM0U=!icon:fa-child!colour:red&marker=latLng:1.2640125925831,103.81227271849899!iwt:JTNDcCUzRTMwJTIwTm92JTIwLSUyMDE3NTVoJTIwdG8lMjAyMTU1aCUyMC0lQzIlQTBTdXNoaSUyMEppcm8lQzIlQTAlM0MlMkZwJTNF!icon:fa-child!colour:red&marker=latLng:1.28886291624463,103.846555999235!iwt:JTNDcCUzRVNRVUUlMjBSb3Rpc3NlcmllJTIwJTI2YW1wJTNCJTIwQWxlaG91c2UlM0MlMkZwJTNF!icon:fa-child!colour:red&marker=latLng:1.28950956232932,103.855665761542!iwt:JTNDcCUzRU1ha2FuJTIwU3V0cmElMjAlNDAlMjBHbHV0dG9ucyUyMEJheSUzQyUyRnAlM0U=!icon:fa-child!colour:red&marker=latLng:1.3492663625637298,103.848643809529!iwt:JTNDcCUzRTI2JTIwTm92JTIwMTgwMGglMjB0byUyMDE5MDBoJTIwLSUyMFMxMSUyMEJpc2hhbiUzQyUyRnAlM0U=!icon:fa-child!colour:red&marker=latLng:1.27948100992947,103.844116761408!iwt:JTNDcCUzRTI1JTIwTm92JTIwLSUyMDE0MDBoJTIwdG8lMjAxNTAwaCUyMC0lMjBkLm8uYyUyMCVFMiU4MCU5MyUyMFRhbmpvbmclMjBQYWdhciUzQyUyRnAlM0U=!icon:fa-child!colour:red&popupWidth=200
