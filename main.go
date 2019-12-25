package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/app"
	"github.com/jboulet/fizzbuzz-go/config"
	"github.com/jboulet/fizzbuzz-go/db"
	"log"
	"net/http"
)

func main() {
	config.UpdateEnv()

	s := config.GetSpecification()

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Decoder:  schema.NewDecoder(),
		Database: database,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":"+s.SERVER_PORT, app.Router))
}
