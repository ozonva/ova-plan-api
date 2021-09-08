package config

import "os"

type envVarKafkaConfig struct {
}

func (c *envVarKafkaConfig) GetBrokers() []string {
	broker := os.Getenv("OVA_KAFKA_BROKER")

	return []string{broker}
}

func NewKafkaConfig() KafkaConfig {
	return &envVarKafkaConfig{}
}
