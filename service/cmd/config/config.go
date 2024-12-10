package config

import (
	"orderAPI/service/pkg/postgres"
	"orderAPI/service/pkg/kafka"
	"os"
	"gopkg.in/yaml.v3"
)

type DBConfig struct {
	DMBS		string 					`yaml:"dmbs"`
    Postgres 	postgres.PostgresConfig	`yaml:"postgres"`
}

type ServerConfig struct {
	Host	string	`yaml:"host"`
	Port	int		`yaml:"port"`
}

type BrokerConfig struct {
	Name	string 				`yaml:"name"`
	Kafka	kafka.KafkaConfig	`yaml:"kafka"`
}

type Config struct {
	DB			DBConfig 		`yaml:"database"`
	Broker		BrokerConfig	`yaml:"broker"`
	Server		ServerConfig	`yaml:"server"`
}

func InitConfig(filePath string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(filePath)
	if err != nil {
	    return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
	    return nil, err
	}
	return &config, nil
}
