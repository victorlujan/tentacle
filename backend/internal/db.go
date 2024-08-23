package internal

import (
	"errors"
	"os"
	"reflect"

	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (c DBConfig) Validate() error {
	v := reflect.ValueOf(c)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).String() == "" {
			return errors.New("missing required DB field: " + v.Type().Field(i).Name)
		}
	}
	return nil

}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func NewDB() (*sqlx.DB, error) {
	config := NewDBConfig()
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return nil, nil
}
