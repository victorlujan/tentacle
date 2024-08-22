package backend

import (
	"context"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"github.com/victorlujan/tentacle/internal"
)

type App struct {
	Ctx     context.Context
	Log     *logrus.Logger
	DB      *gorm.DB
	LogFile string
	DBFile  string
}

func NewApp() *App {
	return &App{}
}

func (a *App) OnStartup(ctx context.Context) {
	a.Ctx = ctx

	a.LogFile = "tentacle.log"
	a.Log = internal.NewLoger(a.LogFile)
	a.Log.Info("Starting Tentacle")

	return
}
