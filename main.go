package main

import (
	"KiasiBot/server"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// wg := new(sync.WaitGroup)

	// add two goroutines
	// wg.Add(2)

	// check for port number
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

	// go routine to launch http server
	// go func() {
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalln("Failed to initialize server at port", port)
	}
	//	wg.Done()
	// }()

	// go routine to launch bot server
	// go func() {
	//	telebot.StartBot()
	//	wg.Done()
	// }()

	// wait until waitgroup is done
	// wg.Wait()
}
