package config

type constKafkaConfig struct {
}

func (c *constKafkaConfig) GetBrokers() []string {
	return []string{"127.0.0.1:9092"}
}

// TODO: Use config
func NewKafkaConfig() *constKafkaConfig {
	return &constKafkaConfig{}
}
