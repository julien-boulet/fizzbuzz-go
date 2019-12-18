package service

import (
	"encoding/json"
	"fmt"
	"github.com/jboulet/fizzbuzz-go/config"
	"github.com/jboulet/fizzbuzz-go/dto"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
	"log"
)

func PushtoKafka(producer *kafka.Producer, gameParameter *dto.GameParameter) {

	gameParameterJson, err := json.Marshal(gameParameter)
	if err != nil {
		log.Println("gameParameter : ", gameParameter)
		log.Fatal("Error convert gameParameter to json : ", err)
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	producer.Produce(&kafka.Message{TopicPartition: kafka.TopicPartition{Topic: &config.GetSpecification().TOPIC, Partition: kafka.PartitionAny}, Value: gameParameterJson}, nil)
}
