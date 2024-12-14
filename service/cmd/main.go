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
	"os"
	"os/signal"
	"syscall"
	"time"
	"log"
)

func main() {
	configPath := "service/configs/conf.yaml"
	config, err := conf.InitConfig(configPath)
	if err != nil {
		log.Fatal("error init config:", err)
	}
	pgConn, err := pkgPostgres.NewConnect(config.DB.Postgres)
	if err != nil {
		log.Fatal("error connecting to DB:", err)
	}
	defer pgConn.Close(context.Background())

	kafkaConsumer, err := pkgKafka.NewConsumer(config.Broker.Kafka)
	if err != nil {
		log.Fatal("error connecting to kafka:", err)
		return
	}
	storage := postgres.New(pgConn)
	cache := cache.New(50)
	repo := repository.New(storage, cache)
	useCase := ucOrder.New(*repo)
	kafkaHandler := kafka.New(kafkaConsumer, useCase)

	log.Println("start processing messages from kafka")
	go kafkaHandler.Start()

	server := http.NewServer(useCase)
	go server.StartServer()

	sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
    log.Println("Received signal:", sig)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := server.Shutdown(ctx); err != nil {
        log.Fatal("Shutdown():", err)
    }
    log.Println("Gracefully shut down")
}