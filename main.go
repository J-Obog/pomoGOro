package main

import (
	"log"
	"net/http"

	"github.com/J-Obog/pomoGOro/api"
	"github.com/J-Obog/pomoGOro/mware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)


func main() {
	//make db migrations
	MakeMigrations()

	//create and configure main router
	router := mux.NewRouter().StrictSlash(true)
	router.Use(mware.ReqLogger)
	router.Use(mware.MimeTypeRes)
	router.Use(handlers.CORS())

	//configure api task routes
	api.AddRoutes(router.PathPrefix("/api").Subrouter())

	//spin up server
	log.Println("Server running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}