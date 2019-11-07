package main

import (
	"github.com/gorilla/schema"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jboulet/fizzbuzz-go/app"
)

func main() {

	app := &app.App{
		Router:  mux.NewRouter().StrictSlash(true),
		Decoder: schema.NewDecoder(),
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
