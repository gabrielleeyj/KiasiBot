package main

import (
	"time"

	"github.com/yanzay/tbot"
)

// Location represents a point on the map
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

func (a *application) startHandler(m *tbot.Message) {
	msg := "This bot has started. Please select a command"
	a.client.SendChatAction(m.Chat.ID, tbot.ActionTyping)
	time.Sleep(1 * time.Second)
	a.client.SendMessage(m.Chat.ID, msg)
}

// func (a *application) InlineQueryHandler(m *tbot.ChosenInlineResult) {

// 	Location := Location{
// 		Longitude: m.Location.Longitude,
// 		Latitude:  m.Location.Latitude,
// 	}
// 	fmt.Println(Location)

// 	a.client.SendMessage(m.InlineMessageID, "Location Received")
// }

// func (a *application) LocationHandler(m *tbot.Message) {
// 	res := "https://www.onemap.sg/amm/amm.html?mapStyle=Default&zoomLevel=15&marker=latLng:1.2843237295033,103.85956355313499!iwt:JTNDcCUzRTI3JTIwTm92JTIwLSUyMDE0MjBoJTIwdG8lMjAxOTAwaCUyMC0lMjBNQlMlMjBDYXNpbm8lM0MlMkZwJTNF!icon:fa-child!colour:red&marker=latLng:1.3020840794266,103.849621716147!iwt:JTNDcCUzRTI1JTIwTm92JTIwLSUyMDE1MzBoJTIwdG8lMjAxODAwaCUyMC0lMjBXaWxraWUlMjBFZGdlJTNDJTJGcCUzRQ==!icon:fa-child!colour:red&marker=latLng:1.35671611225461,103.986514607903!iwt:JTNDcCUzRTI2JTIwTm92JTIwMjA1NWglMjB0byUyMDIxNTVoJTIwLSUyMEtvcGl0aWFtJTNDJTJGcCUzRSUwQSUzQ3AlM0UyNCUyME5vdiUyMDIxMDVoJTIwdG8lMjAyMjAwaCUyMC0lMjBLb3BpdGlhbSUzQyUyRnAlM0U=!icon:fa-child!colour:red&marker=latLng:1.2640125925831,103.81227271849899!iwt:JTNDcCUzRTMwJTIwTm92JTIwLSUyMDE3NTVoJTIwdG8lMjAyMTU1aCUyMC0lQzIlQTBTdXNoaSUyMEppcm8lQzIlQTAlM0MlMkZwJTNF!icon:fa-child!colour:red&marker=latLng:1.28886291624463,103.846555999235!iwt:JTNDcCUzRVNRVUUlMjBSb3Rpc3NlcmllJTIwJTI2YW1wJTNCJTIwQWxlaG91c2UlM0MlMkZwJTNF!icon:fa-child!colour:red&marker=latLng:1.28950956232932,103.855665761542!iwt:JTNDcCUzRU1ha2FuJTIwU3V0cmElMjAlNDAlMjBHbHV0dG9ucyUyMEJheSUzQyUyRnAlM0U=!icon:fa-child!colour:red&marker=latLng:1.3492663625637298,103.848643809529!iwt:JTNDcCUzRTI2JTIwTm92JTIwMTgwMGglMjB0byUyMDE5MDBoJTIwLSUyMFMxMSUyMEJpc2hhbiUzQyUyRnAlM0U=!icon:fa-child!colour:red&marker=latLng:1.27948100992947,103.844116761408!iwt:JTNDcCUzRTI1JTIwTm92JTIwLSUyMDE0MDBoJTIwdG8lMjAxNTAwaCUyMC0lMjBkLm8uYyUyMCVFMiU4MCU5MyUyMFRhbmpvbmclMjBQYWdhciUzQyUyRnAlM0U=!icon:fa-child!colour:red&popupWidth=200"
// 	a.client.SendChatAction(m.Chat.ID, tbot.ActionTyping)
// 	time.Sleep(1 * time.Second)
// 	a.client.SendMessage(m.Chat.ID, res)
// }

func (a *application) LocationHandler(m *tbot.Message) {
	buttons := makeButtons()
	a.client.SendMessage(m.Chat.ID, "Pick an option: ", tbot.OptInlineKeyboardMarkup(buttons))
}

func (a *application) LocTestHandler(m *tbot.Message) {
	button := locButton()
	a.client.SendMessage(m.Chat.ID, "Send Location?", tbot.OptReplyKeyboardMarkup(button))
	// a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	// a.client.SendLocation(m.Chat.ID, 1.352083, 103.819839)
}

func (a *application) callBackHandler(cq *tbot.CallbackQuery) {
	// options := cq.Data
	msg := cq.Data

	a.client.SendMessage(cq.Message.Chat.ID, msg)
}
