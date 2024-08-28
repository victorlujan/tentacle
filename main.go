package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"github.com/victorlujan/tentacle/backend"
	"github.com/victorlujan/tentacle/config"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	app := backend.NewApp()

	err := wails.Run(&options.App{
		Title:  config.Title,
		Width:  config.Width,
		Height: config.Height,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		LogLevel:         logger.DEBUG,
		BackgroundColour: &options.RGBA{R: 255, G: 0, B: 0, A: 128},
		OnStartup:        app.OnStartup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
