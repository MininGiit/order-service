package kafka

import (
	"bytes"
	"encoding/json"
	"log"
	"orderAPI/service/internal/domain/order"
	"orderAPI/service/internal/usecase"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaHandler struct {
	uc			usecase.Order
	consumer	*kafka.Consumer
}

//перенести инициализацию консьюмера в pkg + конфиги

func New(uc usecase.Order) (*KafkaHandler, error) {
	//Добавить парсе конфигов
	config := kafka.ConfigMap{
        "bootstrap.servers": "localhost:9092",
        "group.id":          "myGroup",
        "auto.offset.reset": "earliest",
        "enable.auto.commit": false, // Отключаем автоматическое подтверждение
    }
	consumer, err := kafka.NewConsumer(&config)
	if err != nil {
		return nil, err
	}
	return &KafkaHandler{uc: uc, consumer: consumer}, nil
}

func (k *KafkaHandler) Start() error {
	err := k.consumer.Subscribe("orders", nil)
	if err != nil {
		return err
	}
	k.consumeMessages()
	return nil
}

func (k *KafkaHandler) consumeMessages() {
    for {
        msg, err := k.consumer.ReadMessage(-1)
        if err == nil {
            log.Printf("Received message: %s", string(msg.Value))
			var order order.Order 
			err := json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&order)
			if err != nil {
				log.Println("error encoding message")
			}
			err = k.uc.Save(&order)
			if err != nil {
				log.Printf("Error save message: %v", err)
			}
            _, err = k.consumer.CommitMessage(msg)
            if err != nil {
                log.Printf("Error committing message: %v", err)
            } else {
				log.Printf("Commit message: %s", string(msg.Value))
			}
        } else {
            log.Printf("Error while reading message: %v", err)
        }
    }
}
