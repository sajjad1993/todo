package meesage_broker

import "time"

type Config interface {
	GetAmqpAddress() string
	GetRetryDelayConnect() time.Duration
	GetRetryAttemptsConnect() uint
}
