package meesage_broker

import (
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

// QueueDeclare creates a queue by given key. if it doesn't exist
func (r *Rabbit) QueueDeclare(key string) error {
	_, err := r.channel.QueueDeclare(
		key,
		false,
		false,
		false,
		false,
		nil,
	)
	return err
}

// Consume consumes messages from a given key
func (r *Rabbit) Consume(key string) (<-chan amqp.Delivery, error) {
	messages, err := r.channel.Consume(
		key,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	return messages, err

}

func NewProducer(config Config) (Producer, error) {
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

func NewConsumer(config Config) (Consumer, error) {
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
