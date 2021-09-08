package kafka

import (
	"github.com/Shopify/sarama"
	cfg "github.com/ozonva/ova-plan-api/internal/config"
	"log"
)

type Producer interface {
	Send(messages Messages) error
	Close() error
}

type syncProducer struct {
	producer sarama.SyncProducer
}

func (p *syncProducer) Send(messages Messages) error {
	kafkaMsgs := make([]*sarama.ProducerMessage, 0, len(messages.GetMessages()))

	for _, message := range messages.GetMessages() {
		kafkaMsgs = append(kafkaMsgs, &sarama.ProducerMessage{
			Topic:     string(messages.GetTopic()),
			Partition: -1,
			Value:     sarama.ByteEncoder(message.GetEncoded()),
		})
	}
	errs := p.producer.SendMessages(kafkaMsgs)
	if errs != nil {
		for _, err := range errs.(sarama.ProducerErrors) {
			log.Println("Write to kafka failed: ", err)
		}
	}
	return errs
}

func (p *syncProducer) Close() error {
	return p.producer.Close()
}

func NewSyncProducer(kafkaConfig cfg.KafkaConfig) Producer {
	// For the access log, we are looking for AP semantics, with high throughput.
	// By creating batches of compressed messages, we reduce network I/O at a cost of more latency.
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(kafkaConfig.GetBrokers(), config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}

	return &syncProducer{producer: producer}
}
