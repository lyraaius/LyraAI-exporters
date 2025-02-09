package dal

import (
	"context"
	"github.com/bytedance/sonic"
	gokafka "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/lyraaius/lyraAiX-exporters/pkg/kafka"
)

func AddUserPredictionLog(ctx context.Context, userId, signalId int64, chainName string, choice int32, txHash, userAddr string) error {

	predictionMap := map[string]interface{}{
		"user_id":    userId,
		"signal_id":  signalId,
		"chain_name": chainName,
		"choice":     choice,
		"tx_hash":    txHash,
		"user_addr":  userAddr,
	}

	data, err := sonic.Marshal(predictionMap)
	if err != nil {
		return err
	}

	err = kafka.KafkaClient.Produce(&gokafka.Message{
		TopicPartition: gokafka.TopicPartition{Topic: &kafka.PredictionTopic, Partition: gokafka.PartitionAny},
		Value:          data,
	}, nil)

	if err != nil {
		return err
	}

	return nil
}
