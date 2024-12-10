package main

import (
	"context"
	"orderAPI/service/internal/delivery/http"
	"orderAPI/service/internal/delivery/kafka"
	"orderAPI/service/internal/infrastructure/cache"
	"orderAPI/service/internal/infrastructure/postgres"
	"orderAPI/service/internal/repository"
	ucOrder "orderAPI/service/internal/usecase/order"
	pkgKafka "orderAPI/service/pkg/kafka"
	pkgPostgres "orderAPI/service/pkg/postgres"
	conf "orderAPI/service/cmd/config"
	"log"
	"fmt"
)

func main() {
	configPath := "service/configs/conf.yaml"
	config, err := conf.InitConfig(configPath)
	if err != nil {
		fmt.Println(err)
	}
	pgConn, err := pkgPostgres.NewConnect(config.DB.Postgres)
	if err != nil {
		fmt.Println(err)
		return
	}
	kafkaConsumer, err := pkgKafka.NewConsumer(config.Broker.Kafka)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Start")
	storage := postgres.New(pgConn)
	cache := cache.New(50)
	repo := repository.New(storage, cache)
	useCase := ucOrder.New(*repo)
	kafkaHandler := kafka.New(kafkaConsumer, useCase)

	log.Println("Start kafka")
	go kafkaHandler.Start()

	server := http.NewServer(useCase)
	log.Println("Start server")
	server.StartServer()
	defer pgConn.Close(context.Background())
}