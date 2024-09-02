package db

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/victorlujan/tentacle/backend/models"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/jmoiron/sqlx"
)

func UpdateUserHalls(ctx context.Context, db *sqlx.DB, userHalls models.UserHalls, logger *logrus.Logger) error {

	logger.Println("Updating User Hall Relations")

	var userHallsInserted int = 0
	//ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Panic(err.Error())
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, "ALTER TABLE users_salones AUTO_INCREMENT=1")
	if err != nil {
		return err
	}

	rows, err := db.Query("SELECT id, email FROM user WHERE email IS NOT NULL ORDER BY id")
	if err != nil {
		return err
	}

	users := make(map[string]string)
	for rows.Next() {
		var id string
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			return err
		}
		users[email] = id
	}

	rows, err = db.Query("SELECT id, codigo FROM salones WHERE codigo IS NOT NULL ORDER BY id")
	if err != nil {
		return err
	}

	halls := make(map[string]string)

	for rows.Next() {
		var id string
		var codigo string
		err = rows.Scan(&id, &codigo)
		if err != nil {
			return err
		}
		halls[codigo] = id
	}

	for _, userHall := range userHalls.Body.ReadMultipleResult.ReadMultipleResult.USERSALON_WS {

		userID, userExists := users[userHall.Usuario]
		hallID, hallExists := halls[userHall.Salon]

		if !userExists || !hallExists {
			logger.Info("Hall or user not exits", userHall.Salon, userHall.Usuario)
			runtime.EventsEmit(ctx, "userHallUpdated", "Hall or user not exits", userHall.Salon, userHall.Usuario)
			continue
		}
		_, err = tx.ExecContext(ctx, "INSERT INTO users_salones (user_id, salon_id) VALUES (?,?)", userID, hallID)
		if err != nil {
			logger.Errorf("Error inserting user hall relation: %v %v", userID, hallID)
			runtime.EventsEmit(ctx, "userHallUpdated", "Error inserting user hall relation", userID, hallID)
			return err
		}

		userHallsInserted++
		progress := (float64(userHallsInserted) / float64(len(userHalls.Body.ReadMultipleResult.ReadMultipleResult.USERSALON_WS)) * 100)
		runtime.EventsEmit(ctx, "userHallUpdated", fmt.Sprintf("User Hall relation inserted: %v %v", userID, hallID))
		runtime.EventsEmit(ctx, "progress", progress)

	}

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	logger.Info("User Hall relations  inserted: ", userHallsInserted)

	return nil
}

func DeleteAllRelations(db *sqlx.DB, logger *logrus.Logger) error {

	logger.Println("Deleting all user hall relations")

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Panic(err.Error())
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "TRUNCATE TABLE users_salones")

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	logger.Info("All user hall relations deleted")

	return nil
}
