package service

import (
	"encoding/json"
	"fmt"
	redis2 "github.com/go-redis/redis/v7"
	"github.com/jboulet/fizzbuzz-go/dto"
	"log"
	"net/http"
)

type SaveRedisService struct {
	Client *redis2.Client
}

func (s *SaveRedisService) Push(gameParameter *dto.GameParameter, req *http.Request) {

	gameParameterJson, err := json.Marshal(gameParameter)
	if err != nil {
		log.Println("gameParameter : ", gameParameter)
		log.Fatal("Error convert gameParameter to json : ", err)
	}

	err = s.Client.Set(fmt.Sprintf("address-%s", req.RemoteAddr), gameParameterJson, 0).Err()
	if err != nil {
		log.Fatalln("error when WriteMessages : ", err)
	}
}
