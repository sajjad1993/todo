package meesage_broker

import (
	"github.com/sajjad1993/todo/internal/gateway/config"
	"github.com/streadway/amqp"
)

type Rabbit struct {
	channel *amqp.Channel
}

// Publish publishes any message with json body
func (r *Rabbit) Publish(key string, body []byte) error {
	err := r.channel.Publish(
		"",
		key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		},
	)
	return err
}

// CreateQueue creates a queue by given key.
func (r *Rabbit) CreateQueue(key string) error {
	_, err := r.channel.QueueDeclare(
		"TestQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	return err
}

func New(config config.Config) (Producer, error) {
	conn, err := amqp.Dial(config.GetAmqpAddress())
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	r := Rabbit{
		channel: ch,
	}
	return &r, nil
}
