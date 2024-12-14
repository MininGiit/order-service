package kafka

import (
	"bytes"
	"encoding/json"
	"log"
	"orderAPI/service/internal/domain/order"
	"orderAPI/service/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-playground/validator/v10"
)

type KafkaHandler struct {
	uc			usecase.Order
	consumer	*kafka.Consumer
}

func New(consumer *kafka.Consumer, uc usecase.Order) (*KafkaHandler) {
	return &KafkaHandler{uc: uc, consumer: consumer}
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
	validate := validator.New()
    for {
        msg, err := k.consumer.ReadMessage(-1)
        if err == nil {
			var order order.Order 
			err := json.NewDecoder(bytes.NewReader(msg.Value)).Decode(&order)
			if err != nil {
				log.Println("error encoding message:", err)
				k.consumer.CommitMessage(msg)
				continue
			}
			err = validate.Struct(order)
			if err != nil {
				log.Println("error validate message:", err)
				k.consumer.CommitMessage(msg)
				continue
			}
			err = k.uc.Save(&order)
			if err != nil {
				log.Printf("error save message: %v", err)
				continue
			}
			log.Println("message processed successfully")
            _, err = k.consumer.CommitMessage(msg)
            if err != nil {
                log.Println("error committing message:", err)
            } else {
				log.Println("commit message:")
			}
        } else {
            log.Printf("error while reading message: %v", err)
        }
    }
}
