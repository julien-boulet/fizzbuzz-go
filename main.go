package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/app"
	"github.com/jboulet/fizzbuzz-go/db"
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"net/http"
)

func init() {
	s := config.GetSpecification()

	if err := envconfig.Process("", &s); err != nil {
		log.Fatal("Env variables loading failed: ", err.Error())
	}

	log.Printf("Specification : %v", s)
}

func main() {

	database, err := db.CreateDatabase(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.BD_HOST, s.DB_PORT, s.DB_USERNAME, s.DB_PASSWORD, s.DB_NAME))
	if err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": s.BS_SERVER})
	if err != nil {
		log.Fatal("Kafka producer connection failed: ", err.Error())
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
