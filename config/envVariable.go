package config

type Specification struct {
	BD_HOST        string `default:"localhost"`
	DB_PORT        string `default:"5432"`
	DB_USERNAME    string `default:"postgres"`
	DB_PASSWORD    string `default:"postgres"`
	DB_NAME        string `default:"postgres"`
	SERVER_PORT    string `default:"8080"`
	BS_SERVER      string `default:"localhost:9093"`
	TOPIC          string `default:"myTopic"`
	REDIS_HOST     string `default:"localhost:6379"`
	REDIS_PASSWORD string `default:""`
	REDIS_DB       int    `default:0`
	IS_KAFKA       bool   `default:"false"`
}

var s Specification

func GetSpecification() *Specification {
	return &s
}
