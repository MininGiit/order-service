package kafka

type KafkaConfig struct {
	Host		string	`yaml:"host"`
	Port		int		`yaml:"port"`
	Group		string	`yaml:"group"`
	Reset		string	`yaml:"reset"`	
	AutoCommit	bool	`yaml:"autoCommit"`
}

// host: localhost
// port: 9092
// group: myGroup
// reset: earliest
// autoCommit: true