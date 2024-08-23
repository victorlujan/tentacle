package backend

import (
	"context"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	"github.com/victorlujan/tentacle/backend/internal"
	"github.com/victorlujan/tentacle/backend/models"
	"github.com/victorlujan/tentacle/config"
)

type App struct {
	Ctx     context.Context
	Log     *logrus.Logger
	DB      *sqlx.DB
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

	err := config.Init()
	if err != nil {
		a.Log.Error(err)
	}

	a.Greet("Victor")
	db, err := internal.NewDB()
	if err != nil {
		a.Log.Error(err)
	}

	a.DB = db

	ma := a.GetHalls()
	fmt.Println(ma)

}

func (a *App) Greet(name string) string {
	fmt.Println(os.Getenv("DEBUG"))
	a.Log.Info(os.Getenv("DB_PORT"))
	return "Hello " + name
}

func (a *App) GetHalls() []models.Machine {
	a.Log.Info("Getting Halls")
	var Machine []models.Machine
	err := a.DB.Get(&Machine, "SELECT * FROM machine")
	if err != nil {
		a.Log.Error(err)
	}

	fmt.Print(Machine)

	return Machine

}
