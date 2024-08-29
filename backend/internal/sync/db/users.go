package db

import (
	"context"
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"

	"github.com/victorlujan/tentacle/backend/internal/sync/utils"
	"github.com/victorlujan/tentacle/backend/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

func UpdateUsers(ctx context.Context, db *sqlx.DB, users models.UserEnvelope, logger *logrus.Logger) error {

	logger.Println("Updating users")

	var userUpdated int = 0
	var userInserted int = 0

	//ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err.Error())
	}
	defer tx.Rollback()
	//_, err = tx.ExecContext(ctx, "ALTER TABLE user AUTO_INCREMENT=1")

	var emailUser []string
	db.Select(&emailUser, "SELECT email FROM user")

	for _, user := range users.Body.ReadMultipleResult.ReadMultipleResult.USUARIOSWS {

		var userExists bool = false

		for _, email := range emailUser {
			if email == user.Usuario {
				userExists = true
			}
		}
		password := utils.HashPassword(user.Clave)

		if userExists {

			// _, err = tx.ExecContext(ctx, "UPDATE user SET updated_on = ?, name = ?, username = ?, nif = ?, email = ?, password = ?, delegation = ?,active = ?,is_rotative = ?,employee_code = ? WHERE email = ?",
			// 	time.Now(), user.NombreEmpleado, user.ApellidosEmpleado, user.DNIEmpleado, user.Usuario, password, user.Delegacion, utils.IsActivo(user.Inactivo), utils.IsRotaturnos(user.Rotaturnos), user.CodEmpleado, user.Usuario)
			// if err != nil {
			// 	return err
			// }
			// userUpdated++

			_, err = tx.ExecContext(ctx, "UPDATE user SET updated_on = ?, hash_key = ? WHERE email = ?",
				time.Now(), password, user.Usuario)
			if err != nil {
				return err
			}
			userUpdated++
			logger.Info("User updated: ", user.Usuario)
			runtime.EventsEmit(ctx, "userUpdated", fmt.Sprintf("User updated: %s", user.Usuario))
			progress := (float64(userUpdated+userInserted) / float64(len(users.Body.ReadMultipleResult.ReadMultipleResult.USUARIOSWS))) * 100
			runtime.EventsEmit(ctx, "progress", progress)
		} else {
			// _, err = tx.ExecContext(ctx, "INSERT INTO user (created_on,updated_on, name, username, nif, email, password, delegation,active,is_rotative,employee_code) VALUES (?,?,?,?,?,?,?,?,?,?,?)",
			// 	time.Now(), time.Now(), user.NombreEmpleado, user.ApellidosEmpleado, user.DNIEmpleado, user.Usuario, password, user.Delegacion, utils.IsActivo(user.Inactivo), utils.IsRotaturnos(user.Rotaturnos), user.CodEmpleado)
			// if err != nil {
			// 	return err
			// }
			userInserted++
			logger.Info("User not inserted: ", user.Usuario)

		}
	}

	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	logger.Println("Users updated: ", userUpdated)
	logger.Println("Users inserted: ", userInserted)

	return nil
}

func DeactivateUsers(db *sqlx.DB, logger *logrus.Logger) error {

	logger.Info("Deactivating users")

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "UPDATE user SET active = 0")
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	logger.Info("Users deactivated")

	return nil
}
