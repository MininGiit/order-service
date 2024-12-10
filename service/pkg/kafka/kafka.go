package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
)

func NewConsumer(config KafkaConfig) (*kafka.Consumer, error){
	addres := fmt.Sprintf("%s:%d", config.Host, config.Port)
	kafkaConfig := kafka.ConfigMap{
		"bootstrap.servers": 	addres,
		"group.id":				config.Group,
		"auto.offset.reset": 	config.Reset,
		"enable.auto.commit":	config.AutoCommit,
	}
	consumer, err := kafka.NewConsumer(&kafkaConfig)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}