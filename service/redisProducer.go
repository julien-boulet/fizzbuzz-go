package service

import (
	"encoding/json"
	"fmt"
	redis2 "github.com/go-redis/redis/v7"
	"github.com/jboulet/fizzbuzz-go/dto"
	"log"
	"net/http"
)

func PushtoRedis(client *redis2.Client, gameParameter *dto.GameParameter, req *http.Request) {

	gameParameterJson, err := json.Marshal(gameParameter)
	if err != nil {
		log.Println("gameParameter : ", gameParameter)
		log.Fatal("Error convert gameParameter to json : ", err)
	}

	err = client.Set(fmt.Sprintf("address-%s", req.RemoteAddr), gameParameterJson, 0).Err()
	if err != nil {
		log.Fatalln("error when WriteMessages : ", err)
	}
}
