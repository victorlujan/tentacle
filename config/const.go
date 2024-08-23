package config

import (
	"github.com/joho/godotenv"
)

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

func Init() {
	godotenv.Load()
}
