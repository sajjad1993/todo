package meesage_broker

type Producer interface {
	Publish(key string, body []byte) error
	QueueDeclare(key string) error
}
