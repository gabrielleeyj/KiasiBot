package main

import (
	"KiasiBot/server"
	"KiasiBot/telebot"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
)

// initialize check for .env file existence
func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}

func main() {

	// create waitgroup
	wg := new(sync.WaitGroup)

	// add 2 go routines to wait group
	wg.Add(2)

	// create router handler
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	pr := server.NewPresenter()
	r.Get("/", pr.Home())
	r.Get("/css/*", pr.CSS("./static/css"))
	r.Get("/js/*", pr.JavaScript("./static/js"))
	r.Get("/json/*", pr.JSON("./static/json"))
	log.Println("Initializing server at port 8080")

	// go routine to launch bot server
	go func() {
		telebot.StartBot()
		wg.Done() // one goroutine finished
	}()

	// go routine to launch http server
	go func() {
		err := http.ListenAndServe(":8080", r)
		if err != nil {
			log.Fatalln("Failed to initialize server at port 8080")
		}
		wg.Done() // second goroutine finished
	}()

	// wait until WaitGroup is done
	wg.Wait()
}
