package config

import (
	"fmt"
	"github.com/sajjad1993/todo/pkg/db"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/spf13/viper"
	"time"
)

type Config interface {
	GetAmqpAddress() string
	GetDatabaseDsn() string
	GetDbDriver() string
	GetUserServiceAddress() string
	GetToDoServiceAddress() string
	GetRetryDelayConnect() time.Duration
	GetRetryAttemptsConnect() uint
}
type SampleConfig struct {
	AmqpAddress          string        `mapstructure:"AMQP_ADDRESS"`
	DatabaseDsn          string        `mapstructure:"DATABASE_DSN"`
	DBDriver             string        `mapstructure:"DATABASE_DRIVER"`
	UserServiceAddress   string        `mapstructure:"USER_SERVICE_ADDRESS"`
	ToDoServiceAddress   string        `mapstructure:"TODO_SERVICE_ADDRESS"`
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

func NewDatabaseConfig(config Config) db.Config {
	return config
}
func NewMessageBrokerConfig(config Config) meesage_broker.Config {
	return config
}

func (s SampleConfig) GetAmqpAddress() string {
	return s.AmqpAddress
}

func (s SampleConfig) GetDatabaseDsn() string {
	return s.DatabaseDsn
}
func (s SampleConfig) GetDbDriver() string {
	return s.DBDriver
}
func (s SampleConfig) GetUserServiceAddress() string {
	return s.UserServiceAddress
}
func (s SampleConfig) GetToDoServiceAddress() string {
	return s.ToDoServiceAddress
}
func (s SampleConfig) GetRetryDelayConnect() time.Duration {
	return s.RetryDelayConnect
}
func (s SampleConfig) GetRetryAttemptsConnect() uint {
	return s.RetryAttemptsConnect
}
