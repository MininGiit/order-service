package main

import (
	"context"
	"orderAPI/service/internal/delivery/http"
	"orderAPI/service/internal/delivery/kafka"
	"orderAPI/service/internal/infrastructure/cache"
	"orderAPI/service/internal/infrastructure/postgres"
	"orderAPI/service/internal/repository"
	ucOrder "orderAPI/service/internal/usecase/order"

	//	"orderAPI/service/internal/repository/storage/postgres"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
)

func main() {
	//Вынести подключение в pkg и добавить конфиги
	connString := "postgres://myuser:mypassword@localhost:5436/mydb?sslmode=disable"

	// Подключение к базе данных
	log.Println("Connect to db")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
	 	log.Fatal(err)
	}

	log.Println("Start")
	storage := postgres.New(conn)
	cache := cache.New(50)
	repo := repository.New(storage, cache)
	useCase := ucOrder.New(*repo)
	kafkaHandler, _ := kafka.New(useCase)

	log.Println("Start kafka")
	go kafkaHandler.Start()

	server := http.NewServer(useCase)
	log.Println("Start server")
	server.StartServer()
	defer conn.Close(context.Background())
}