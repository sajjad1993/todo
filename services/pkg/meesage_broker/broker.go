package meesage_broker

import "github.com/streadway/amqp"

type Broker interface {
	Producer
	Consumer
}
type Producer interface {
	Publish(key string, body []byte) error
	QueueDeclare(key string) error
}

type Consumer interface {
	Consume(key string) (<-chan amqp.Delivery, error)
	QueueDeclare(key string) error
}
