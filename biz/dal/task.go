package dal

import (
	"context"
	"github.com/bytedance/sonic"
	gokafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lyraaius/lyraAiX-exporters/pkg/kafka"
	"time"
)

func DailySignIn(ctx context.Context, userID int64, taskID int32, chainName string) (success bool, err error) {

	predictionMap := map[string]interface{}{
		"user_id":    userID,
		"task_id":    taskID,
		"chain_name": chainName,
		"timestamp":  time.Now().Unix(),
	}

	data, err := sonic.Marshal(predictionMap)
	if err != nil {
		return false, err
	}

	err = kafka.KafkaClient.Produce(&gokafka.Message{
		TopicPartition: gokafka.TopicPartition{Topic: &kafka.CheckInTopic, Partition: gokafka.PartitionAny},
		Value:          data,
	}, nil)

	if err != nil {
		return false, err
	}

	return true, nil
}
