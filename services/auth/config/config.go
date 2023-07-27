package config

import (
	"fmt"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/spf13/viper"
	"time"
)

type Config interface {
	GetAmqpAddress() string
	GetUserServiceAddress() string
	GetAuthServiceAddress() string
	GetAccessSignKey() string
	GetAccessJWTExp() time.Duration
	GetRetryDelayConnect() time.Duration
	GetRetryAttemptsConnect() uint
}
type SampleConfig struct {
	AmqpAddress        string `mapstructure:"AMQP_ADDRESS"`
	UserServiceAddress string `mapstructure:"USER_SERVICE_ADDRESS"`
	AuthServiceAddress string `mapstructure:"AUTH_SERVICE_ADDRESS"`

	AccessSignKey        string        `mapstructure:"ACCESS_SIGN_KEY"`
	AccessJWTExp         time.Duration `mapstructure:"ACCESS_JWT_EXP" `
	RetryDelayConnect    time.Duration `mapstructure:"RETRY_DELAY_CONNECT"`
	RetryAttemptsConnect uint          `mapstructure:"RETRY_ATTEMPTS_CONNECT" `
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

func (s SampleConfig) GetAccessSignKey() string {
	return s.AccessSignKey
}
func (s SampleConfig) GetAccessJWTExp() time.Duration {
	return s.AccessJWTExp
}
func (s SampleConfig) GetUserServiceAddress() string {
	return s.UserServiceAddress
}
func (s SampleConfig) GetAuthServiceAddress() string {
	return s.AuthServiceAddress
}
func (s SampleConfig) GetRetryDelayConnect() time.Duration {
	return s.RetryDelayConnect
}
func (s SampleConfig) GetRetryAttemptsConnect() uint {
	return s.RetryAttemptsConnect
}
