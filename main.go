package main

import (
	"KiasiBot/server"
	"KiasiBot/telebot"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// // initialize check for .env file existence
// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error Loading .env file")
// 	}
// }

// var (
// 	// DBURI API token
// 	DBURI = os.Getenv("DB_URI")
// 	// TelegramToken API TOKEN
// 	TelegramToken = os.Getenv("TELEGRAM_TOKEN")
// )

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// create router handler
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	pr := server.NewPresenter()
	r.Get("/", pr.Home())
	r.Get("/css/*", pr.CSS("./static/css"))
	r.Get("/js/*", pr.JavaScript("./static/js"))
	r.Get("/json/*", pr.JSON("./static/json"))
	log.Println("Initializing server at port :", port)

	// go routine to launch bot server
	go func() {
		telebot.StartBot()

	}()

	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalln("Failed to initialize server at port", port)
	}

}
