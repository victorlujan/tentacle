package backend

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/jmoiron/sqlx"
	"github.com/victorlujan/tentacle/backend/internal"
	"github.com/victorlujan/tentacle/backend/internal/sync/db"
	"github.com/victorlujan/tentacle/backend/internal/sync/services"
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

	db, err := internal.NewDB()
	if err != nil {
		a.Log.Error(err)
	}
	a.DB = db
}

func (a *App) Greet(name string) string {
	return "Hello " + name
}

func (a *App) GetMachines() []models.Machine {
	a.Log.Info("Getting Halls")
	var Machine []models.Machine
	err := a.DB.Select(&Machine, "SELECT id, description FROM machine")
	if err != nil {
		a.Log.Error(err)
	}

	return Machine
}

func (a *App) GetUsers() []models.User {
	a.Log.Info("Getting Users")
	var User []models.User
	err := a.DB.Select(&User, "SELECT id, COALESCE(email, 'Desconocido') as email, COALESCE(nif, 'Desconocido') as nif , COALESCE(delegation, 'Desconocido') as delegation  FROM user;")
	if err != nil {
		a.Log.Error(err)
	}

	return User
}

func (a *App) SyncUsers() bool {
	a.Log.Info("Syncing users")
	var initTime time.Time = time.Now()

	users, err := services.GetUsers()
	if err != nil {
		a.Log.Error(err)
		runtime.EventsEmit(a.Ctx, "userUpdated", err.Error())

		return false
	}

	err = db.UpdateUsers(a.Ctx, a.DB, users, a.Log)
	if err != nil {
		a.Log.Error(err)
		runtime.EventsEmit(a.Ctx, "userUpdated", err.Error())
		return false
	}

	var endTime time.Time = time.Now()
	var duration time.Duration = endTime.Sub(initTime)
	a.Log.Info("Sync users took: ", duration)

	return true
}

func (a *App) SyncHalls() bool {
	a.Log.Info("Syncing halls")
	var initTime time.Time = time.Now()

	halls, err := services.GetHalls()
	if err != nil {
		a.Log.Error(err)

		return false
	}

	err = db.UpdateHalls(a.Ctx, a.DB, halls, a.Log)
	if err != nil {
		a.Log.Error(err)
		return false
	}

	var endTime time.Time = time.Now()
	var duration time.Duration = endTime.Sub(initTime)
	a.Log.Info("Sync halls took: ", duration)

	return true
}

func (a *App) SyncUserHalls() bool {
	a.Log.Info("Syncing user halls")
	var initTime time.Time = time.Now()

	userHalls, err := services.GetUserHalls()
	if err != nil {
		a.Log.Error(err)
		return false
	}

	err = db.DeleteAllRelations(a.DB, a.Log)
	if err != nil {
		a.Log.Error(err)
		return false
	}

	err = db.UpdateUserHalls(a.Ctx, a.DB, userHalls, a.Log)
	if err != nil {
		a.Log.Error(err)
		return false
	}

	var endTime time.Time = time.Now()
	var duration time.Duration = endTime.Sub(initTime)
	a.Log.Info("Sync user halls took: ", duration)

	return true
}

func (a *App) EmitTestEvent() {
	runtime.EventsEmit(a.Ctx, "test_event", "test_data")

}

func (a *App) LogEmiter(eventName string, data string) {
	a.Log.Info("Emiting event: ", eventName)
	runtime.EventsEmit(a.Ctx, eventName, data)
}
