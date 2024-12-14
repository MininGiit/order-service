package tests

import (
	repoPostgres "orderAPI/service/internal/infrastructure/postgres"
	"orderAPI/service/pkg/postgres"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	config := postgres.PostgresConfig{
		Host:		"localhost",
		Port:		5436,
		User:		"myuser", 
		Password: 	"mypassword",
		DBName:   	"mydb",
		SSLMode:	"disable",
	}
	conn, err := postgres.NewConnect(config)
	if err != nil {
		t.Error(err)
	} 
	repo := repoPostgres.New(conn)
	
	order1 := generateOrder()
	repo.Save(order1)
	orderFromDB, err := repo.GetByUID(order1.OrderUID)
	if err != nil || orderFromDB == nil {
		t.Error(err)
	} 
	if !CompareOrders(order1, orderFromDB) {
		t.Error("orders differ")
	}
}