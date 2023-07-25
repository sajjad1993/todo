package config

import (
	"fmt"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/spf13/viper"
	"time"
)

type Config interface {
	GetAmqpAddress() string
	GetHTTPServerReadTimeout() time.Duration
	GetHTTPServerReadHeaderTimeout() time.Duration
	GetHTTPServerWriteTimeout() time.Duration
	GetHTTPServerAddress() string
}
type SampleConfig struct {
	AmqpAddress                 string        `mapstructure:"AMQP_ADDRESS"`
	HTTPServerReadTimeout       time.Duration `mapstructure:"HTTP_SERVER_READ_TIMEOUT"`
	HTTPServerReadHeaderTimeout time.Duration `mapstructure:"HTTP_SERVER_READ_HEADER_TIMEOUT"`
	HTTPServerWriteTimeout      time.Duration `mapstructure:"HTTP_SERVER_WRITE_TIMEOUT"`
	HTTPServerAddress           string        `mapstructure:"HTTP_SERVER_ADDRESS"`
}

func New() (Config, error) {
	viper.SetConfigName("app")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf(" config file: %w", err))
		viper.AutomaticEnv()
	}
	var c SampleConfig
	err = viper.Unmarshal(&c)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	return &c, nil
}
func NewMessageBrokerConfig(config Config) meesage_broker.Config {
	return config
}

func (s SampleConfig) GetAmqpAddress() string {
	return s.AmqpAddress
}
func (s SampleConfig) GetHTTPServerReadTimeout() time.Duration {
	return s.HTTPServerReadTimeout
}
func (s SampleConfig) GetHTTPServerReadHeaderTimeout() time.Duration {
	return s.HTTPServerReadHeaderTimeout
}
func (s SampleConfig) GetHTTPServerWriteTimeout() time.Duration {
	return s.HTTPServerWriteTimeout
}
func (s SampleConfig) GetHTTPServerAddress() string {
	return s.HTTPServerAddress
}
