package postgres

import (
	"fmt"
	"strings"
)

type PostgresConfig struct {
    Host     string	`yaml:"host"`
    Port     int    `yaml:"port"`
    User     string `yaml:"user"`
    Password string `yaml:"password"`
    DBName   string `yaml:"dbname"`
    SSLMode  string `yaml:"sslmode"`
}

// type DBConfig struct {
// 	DMBS		string 			`yaml:"dmbs"`
//     Postgres 	PostgresConfig	`yaml:"postgres"`
// }

// type Config struct {
// 	DataBase	DBConfig `yaml:"database"`
// }

// func InitConfig(filePath string) (*Config, error) {
//     var config Config
//     data, err := os.ReadFile(filePath)
//     if err != nil {
//         return nil, err
//     }

//     err = yaml.Unmarshal(data, &config)
//     if err != nil {
//         return nil, err
//     }
//     return &config, nil
// }

func (s PostgresConfig) toDSN() string {
	var args []string

	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	}

	if s.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	}

	if len(s.DBName) > 0 {
		args = append(args, fmt.Sprintf("dbname=%s", s.DBName))
	}

	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", s.User))
	}

	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	}

	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	}
	return strings.Join(args, " ")
}
