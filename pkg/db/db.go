package db

import (
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"

	"gorm.io/gorm/logger"
	"log"
	"os"

	"gorm.io/gorm"
)

type Config interface {
	GetDatabaseDsn() string
	GetDbDriver() string
}

func NewDb(cfg Config) (*gorm.DB, error) {
	_ = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
	dbConfig, err := pgx.ParseConfig(cfg.GetDatabaseDsn())
	if err != nil {
		return nil, err
	}
	if err != nil {
		log.Println("DB Connection error : ", err.Error())
		return nil, err
	}
	dbSql := stdlib.OpenDB(*dbConfig)
	gdb, err := gorm.Open(postgres.New(postgres.Config{
		Conn: dbSql,
	}), &gorm.Config{})
	gdb.Debug()
	if err != nil {
		return nil, err
	}

	return gdb, nil
}
