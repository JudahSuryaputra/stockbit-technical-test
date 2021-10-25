package repository

import (
	"fmt"
	"log"
	"stockbit-backend/cfg"

	"github.com/gocraft/dbr"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func GetPostgresDSN() (dsn string) {
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		viper.GetString(cfg.PostgresHost),
		viper.GetString(cfg.PostgresUsername),
		viper.GetString(cfg.PostgresPassword),
		viper.GetString(cfg.PostgresDatabaseName))
	log.Printf("%s", dsn)
	return
}

func Conn() (*dbr.Connection, error) {
	conn, err := dbr.Open("postgres", GetPostgresDSN(), nil)

	return conn, err
}
