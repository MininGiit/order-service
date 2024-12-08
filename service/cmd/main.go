package main

import (
	"context"
	"orderAPI/service/internal/delivery/kafka"
	"orderAPI/service/internal/repository/cache/mapCache"
	"orderAPI/service/internal/repository/storage/postgres"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, connString)
	if err != nil {
	 log.Fatal(err)
	}

	storage := postgres.New(conn)
	cache := mapCache.New(50)
	useCase := ucOrder.New(storage, cache)
	kafkaHandler, err := kafka.New(useCase)
	kafkaHandler.Start()

	defer conn.Close(context.Background())
}