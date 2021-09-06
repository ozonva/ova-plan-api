package kafka

type Topic string

type Message interface {
	GetEncoded() []byte
}

type Messages interface {
	GetTopic() Topic
	GetMessages() []Message
}
