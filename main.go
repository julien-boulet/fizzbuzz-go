package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/app"
	"github.com/jboulet/fizzbuzz-go/config"
	"github.com/jboulet/fizzbuzz-go/db"
	"github.com/jboulet/fizzbuzz-go/service"
	"github.com/kelseyhightower/envconfig"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

func init() {
	s := config.GetSpecification()

	if err := envconfig.Process("", s); err != nil {
		log.Fatal("Env variables loading failed: ", err.Error())
	}

	log.Printf("Specification : %v", *s)
}

// @title FizzBuzz Go API
// @version 1.0
// @description This is a simple API that plays the fizzbuzz game and store statistics.

// @schemes http
// @BasePath /
func main() {
	s := config.GetSpecification()

	database, err := db.CreateDatabase(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.BD_HOST, s.DB_PORT, s.DB_USERNAME, s.DB_PASSWORD, s.DB_NAME))
	if err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	}
	defer database.Close()

	app := &app.App{
		Router:   mux.NewRouter().StrictSlash(true),
		Decoder:  schema.NewDecoder(),
		Database: database,
	}

	if s.IS_KAFKA {
		client := kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{s.BS_SERVER},
			Topic:    s.TOPIC,
			Balancer: &kafka.LeastBytes{},
		})
		app.Service = &service.SaveKafkaService{Client: client}
		defer client.Close()
	} else {
		client := redis.NewClient(&redis.Options{
			Addr:     s.REDIS_HOST,
			Password: s.REDIS_PASSWORD, // no password set
			DB:       s.REDIS_DB,       // use default DB
		})
		app.Service = &service.SaveRedisService{Client: client}
		defer client.Close()
	}

	app.SetupRouter()

	log.Fatal(http.ListenAndServe(":"+s.SERVER_PORT, app.Router))
}
