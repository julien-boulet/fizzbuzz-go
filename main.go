package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/app"
	"github.com/jboulet/fizzbuzz-go/config"
	"github.com/jboulet/fizzbuzz-go/db"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
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

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": s.BS_SERVER})
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Decoder:  schema.NewDecoder(),
		Database: database,
		Producer: producer,
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":"+s.SERVER_PORT, app.Router))
}
