package meesage_broker

import "github.com/streadway/amqp"

type Consumer interface {
	Consume(key string) (<-chan amqp.Delivery, error)
}
