package postgres

import (
	"orderAPI/service/internal/domain/order"
	"github.com/jackc/pgx/v4"
	"context"
	"encoding/json"
	"log"
)

type Storage struct {
	db	*pgx.Conn
}

func New(conn *pgx.Conn) *Storage {
	return &Storage{
		db: conn,
	}
}

func (s *Storage) GetByUID(uid string) (*order.Order, error) {
	query := `SELECT order_data FROM orders WHERE order_uid = $1`
	var orderJSON []byte
	s.db.QueryRow(context.Background(), query, uid).Scan(&orderJSON)
	var order order.Order
	err := json.Unmarshal(orderJSON, &order)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &order, nil
}

func (s *Storage) GetAll() ([]*order.Order, error) {
	return nil, nil
}

func (s *Storage) Save(order *order.Order) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		return err
	}
	query := `INSERT INTO orders (order_uid, order_data) VALUES($1, $2)`
	_, err = s.db.Exec(context.Background(), query, order.OrderUID, orderJSON)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
