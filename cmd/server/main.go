package main

import (
	"KiasiBot/server"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	pr := server.NewPresenter()
	r.Get("/", pr.Home())
	r.Get("/css/*", pr.CSS("./static/css"))
	r.Get("/js/*", pr.JavaScript("./static/js"))
	r.Get("/json/*", pr.JSON("./static/json"))
	log.Println("Initializing server at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Failed to initialize server at port 8080")
	}
	return
}
