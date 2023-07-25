package meesage_broker

type Config interface {
	GetAmqpAddress() string
}
