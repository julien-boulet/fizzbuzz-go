package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
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

var S Specification

func UpdateEnv() {
	err := envconfig.Process("", &S)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Specification : ", S)
}
