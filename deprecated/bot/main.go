package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Create a struct that mimics the webhook response body
// https://core.telegram.org/bots/api#update
type webhookReqBody struct {
	Message struct {
		Text     string `json:"text"`
		Location struct {
			longitude float64 `json:"longitude"`
			latitude  float64 `json:"latitude"`
		}
		Chat struct {
			ID       int64  `json:"id"`
			Username string `json:"username"`
		} `json:"chat"`
	} `json:"message"`
}

// Handler is called everytime telegram sends us a webhook event
func Handler(res http.ResponseWriter, req *http.Request) {
	// First, decode the JSON response body
	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	// Check if the message contains the word "marco"
	// if not, return without doing anything
	if !strings.Contains(strings.ToLower(body.Message.Text), "/help") {
		return
	}

	// If the text contains marco, call the `sayPolo` function, which
	// is defined below
	if err := helperFunc(body.Message.Chat.ID); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	// log a confirmation message if the message is sent successfully
	fmt.Println("reply sent")
}

//The below code deals with the process of sending a response message
// to the user

// Create a struct to conform to the JSON body
// of the send message request
// https://core.telegram.org/bots/api#sendmessage
type sendMessageReqBody struct {
	ChatID int64  `json:"chatID"`
	Text   string `json:"text"`
}

// rtnLoc takes a chatID and sends loc data back
func rtnLoc(chatID int64) error {
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   GetLoc(), // sends GetLoc()
	}
}

// sayPolo takes a chatID and sends "polo" to them
func helperFunc(chatID int64) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   "Commands: /checkin , /shareloc, /getloc",
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	// Send a post request with your token
	res, err := http.Post(os.Getenv("TELEGRAM_BOT"), "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

// Main function to start the server on port 3000
func main() {
	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}