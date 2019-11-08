package main

import (
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jboulet/fizzbuzz-go/app"
	"github.com/jboulet/fizzbuzz-go/db"
)

func main() {
	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: %s", err.Error())
	}

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Decoder:  schema.NewDecoder(),
		Database: database,
	}

	app.SetupRouter()

	port := utils.EnvVariable("SERVER_PORT", utils.ServerPort)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}
