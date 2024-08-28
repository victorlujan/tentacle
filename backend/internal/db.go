package internal

import (
	"context"
	"errors"
	"os"
	"reflect"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func NewDB() (*sqlx.DB, error) {
	config := NewDBConfig()
	if err := config.Validate(); err != nil {
		return nil, err
	}

	var url string = config.User + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.DBName + "?parseTime=true"

	ctxTimeout, ctxCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxCancel()

	db, err := sqlx.ConnectContext(ctxTimeout, "mysql", url)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(50)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.PingContext(ctxTimeout)
	if err != nil {
		return nil, err
	}
	return db, nil
}
