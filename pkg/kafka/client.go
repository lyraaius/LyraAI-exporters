package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/deagentAI/alphax-exporters/conf"
)

var (
	CheckInTopic    string
	PredictionTopic string
	KafkaClient     *kafka.Producer
)

func Init(config *conf.Config) {
	var err error

	KafkaClient, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": config.Kafka.ServerAddr,
	})

	if err != nil {
		panic(err)
	}

	CheckInTopic = config.Kafka.CheckInTopic
	PredictionTopic = config.Kafka.PredictionTopic
}
