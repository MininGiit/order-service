package postgres

import (
	"github.com/jackc/pgx/v4"
	"context"
)


func NewConnect(config PostgresConfig) (*pgx.Conn, error){
	connConfig, err := pgx.ParseConfig(config.toDSN()) 
	if err != nil {
		return nil, err
	}
	conn, err := pgx.ConnectConfig(context.Background(), connConfig)
	if err != nil {
		return nil, err
	}
	return conn, nil
}