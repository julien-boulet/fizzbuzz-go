package service

import (
	"encoding/json"
	"fmt"
	"github.com/jboulet/fizzbuzz-go/dto"
	kafka "github.com/segmentio/kafka-go"
	"log"
	"net/http"
)

type SaveKafkaService struct {
	Client *kafka.Writer
}

func (s *SaveKafkaService) Push(gameParameter *dto.GameParameter, req *http.Request) {

	gameParameterJson, err := json.Marshal(gameParameter)
	if err != nil {
		log.Println("gameParameter : ", gameParameter)
		log.Fatal("Error convert gameParameter to json : ", err)
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprintf("address-%s", req.RemoteAddr)),
		Value: gameParameterJson,
	}
	if err := s.Client.WriteMessages(req.Context(), msg); err != nil {
		log.Fatalln("error when WriteMessages : ", err)
	}
}
