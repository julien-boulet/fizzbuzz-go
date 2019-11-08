package utils

import (
	"log"
	"os"
)

var (
	Host       = "localhost"
	Port       = "5432"
	User       = "postgres"
	Password   = "postgres"
	DBName     = "postgres"
	ServerPort = "8080"
)

func UpdateEnv() {
	update("SERVER_PORT", &ServerPort)
	update("BD_HOST", &Host)
	update("DB_PORT", &Port)
	update("DB_USERNAME", &User)
	update("DB_PASSWORD", &Password)
	update("DB_NAME", &DBName)

	log.Println("BD_HOST : ", Host)
}

func update(key string, defaultValue *string) {
	value, ok := os.LookupEnv(key)
	if ok {
		*defaultValue = value
	}
}
