package config

import (
	_ "embed"
	"os"

	"github.com/joho/godotenv"
)

//go:embed config.env
var configEnv string

const (
	App     = "Tentacle"
	Version = "v0.1.0"
)

const (
	Title = App + " " + Version
)

const (
	Width  = 1024
	Height = 768
)

func Init() error {
	env, err := godotenv.Unmarshal(configEnv)
	if err != nil {
		return err
	}
	for key, value := range env {
		err = os.Setenv(key, value)
		if err != nil {
			return err
		}

	}
	return nil

}
