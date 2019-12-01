package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jboulet/fizzbuzz-go/app"
	"github.com/jboulet/fizzbuzz-go/db"
	"github.com/jboulet/fizzbuzz-go/utils"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
	"net/http"
)

func main() {
	utils.UpdateEnv()

	database, err := db.CreateDatabase()
	if err != nil {
		log.Fatal("Database connection failed: ", err.Error())
	}

	//props.put(ProducerConfig.BOOTSTRAP_SERVERS_CONFIG, applicationProperties.getBootstrapServersConfig());
	//props.put(client.id, applicationProperties.getClientIdConfig());
	//props.put(acks, applicationProperties.getAcksConfig());
	//props.put(key.serializer, StringSerializer.class);
	//props.put(value.serializer, StringSerializer.class);

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": utils.BSServers})
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

	log.Fatal(http.ListenAndServe(":"+utils.ServerPort, app.Router))
}
