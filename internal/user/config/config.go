package config

import (
	"fmt"
	"github.com/sajjad1993/todo/pkg/db"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/spf13/viper"
)

type Config interface {
	GetAmqpAddress() string
	GetDatabaseDsn() string
	GetUserServiceAddress() string
}
type SampleConfig struct {
	AmqpAddress        string `mapstructure:"AMQP_ADDRESS"`
	DatabaseDsn        string `mapstructure:"DATABASE_DSN"`
	UserServiceAddress string `mapstructure:"USER_SERVICE_ADDRESS"`
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
func (s SampleConfig) GetUserServiceAddress() string {
	return s.UserServiceAddress
}
