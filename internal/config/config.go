package config

type DatabaseConfig interface {
	GetDsn() string
}

type KafkaConfig interface {
	GetBrokers() []string
}
