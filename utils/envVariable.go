package utils

import "os"

const (
	Host       = "localhost"
	Port       = "5432"
	User       = "postgres"
	Password   = "postgres"
	DBName     = "postgres"
	ServerPort = "8080"
)

func EnvVariable(key string, defaultValue string) string {
	dbHost := defaultValue
	value, ok := os.LookupEnv(key)
	if ok {
		dbHost = value
	}
	return dbHost
}
