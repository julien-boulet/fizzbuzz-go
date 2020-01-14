package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/app"
	"github.com/jboulet/fizzbuzz-go/db"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

type Specification struct {
	BD_HOST     string `default:"localhost"`
	DB_PORT     string `default:"5432"`
	DB_USERNAME string `default:"postgres"`
	DB_PASSWORD string `default:"postgres"`
	DB_NAME     string `default:"postgres"`
	SERVER_PORT string `default:"8080"`
	BS_SERVER   string `default:"localhost:9095,localhost:9093,localhost:9094"`
	TOPIC       string `default:"myTopic"`
}

var s Specification

func init() {
	err := envconfig.Process("", &s)
	if err != nil {
		log.Fatal("Env variables loading failed: ", err.Error())
	}

	log.Printf("Specification : ", s)

}

func main() {

	database, err := db.CreateDatabase(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.BD_HOST, s.DB_PORT, s.DB_USERNAME, s.DB_PASSWORD, s.DB_NAME))
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
