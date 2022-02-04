package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type service struct{}

type KafkaService interface {
	Producer(topic string, message interface{})
}

func NewKafkaService() KafkaService {
	return &service{}
}

func (s *service) Producer(topic string, message interface{}) {
	writer := &kafka.Writer{
		Addr:  kafka.TCP(os.Getenv("KAFKA_BROKER_ADDRESS")),
		Topic: topic,
	}

	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Fatal("error: ", err)
	}

	err = writer.WriteMessages(context.Background(), kafka.Message{
		Value: messageBytes,
	})

	if err != nil {
		log.Fatal("cannot write a message: ", err)
	}
}
