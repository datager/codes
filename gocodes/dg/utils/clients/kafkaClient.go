package clients

import (
	"io"
)

type KafkaProducer interface {
	io.Closer
	Send(data []byte) error
}

type KafkaConsumer interface {
	io.Closer
	GetMessage() ([]byte, error)
}
