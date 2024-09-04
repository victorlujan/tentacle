package db

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/victorlujan/tentacle/backend/internal/sync/utils"
	"github.com/victorlujan/tentacle/backend/models"
	"github.com/wailsapp/wails/v2/pkg/runtime"

	"log/slog"

	"github.com/jmoiron/sqlx"
)

func UpdateHalls(ctx context.Context, db *sqlx.DB, halls models.Halls, logger *logrus.Logger) error {

	logger.Info("Updating Halls")
	var hallsUpdated int = 0
	var hallsInserted int = 0
	//ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		logger.Panic(err.Error())
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, "ALTER TABLE salones AUTO_INCREMENT=1")
	if err != nil {
		return err
	}
	var hallCode []string
	db.Select(&hallCode, "SELECT codigo FROM salones where codigo is not NULL")

	for _, hall := range halls.Body.ReadMultipleResult.ReadMultipleResult.CONFSALONES_WS {
		var hallExists bool = false

		for _, code := range hallCode {
			if code == hall.CodigoSalon {
				hallExists = true
			}
		}
		if hallExists {
			region := utils.NewNullString(hall.RegionCode)
			_, err = tx.ExecContext(ctx, "UPDATE salones SET updated_on =?, nombre = ?, region_code = ?, direccion = ?, localidad = ?, cif = ?, empresa = ?, mail_responsable = ?, mail_coordinador = ?, sindni = ?, path_music = ?, is_withholding =?, is_region_control = ?, is_hall = ?, telefono = ?, hora_apertura = ? ,hora_cierre = ?, activo = ? , is_invitation = ?, external =? WHERE codigo = ?",
				time.Now(), hall.Nombre, region, hall.Direccion, hall.Localidad, hall.Cif, hall.Empresa, hall.Correoresponsable, hall.CorreoCoordinador, utils.IsSinDNI(hall.SinDNI), hall.RutaMusica, utils.Socio(hall.SalonSocio), hall.ExcluirControlAcceso, utils.IsHall(hall.IsHall), hall.Telefono, hall.HoraApertura, hall.HoraCierre, utils.IsActive(hall.ActivoEmotivanet), utils.IsInvitation(hall.InvitacionesSinCliente), utils.IsExternal(hall.Externo), hall.CodigoSalon)
			if err != nil {
				return err
			}

			hallsUpdated++
			runtime.EventsEmit(ctx, "hallUpdated", fmt.Sprintf("Hall updated: %s", hall.CodigoSalon))

		} else {
			_, err = tx.ExecContext(ctx, "INSERT INTO salones (created_on,updated_on,nombre, cif, empresa,direccion,codigo, localidad,secuencia, mail_responsable, mail_coordinador,sindni, ip_salon, activo, hora_apertura, hora_cierre, path_music, is_invitation,is_withholding,is_region_control,region_code,is_hall,external) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
				time.Now(), time.Now(), hall.Nombre, hall.Cif, hall.Empresa, hall.Direccion, hall.CodigoSalon,
				hall.Localidad, 0, hall.Correoresponsable, hall.CorreoCoordinador, utils.IsSinDNI(hall.SinDNI),
				utils.NewNullString(hall.DireccionIP), utils.IsActive(hall.ActivoEmotivanet), hall.HoraApertura, hall.HoraCierre,
				hall.RutaMusica, utils.IsInvitation(hall.InvitacionesSinCliente), utils.Socio(hall.SalonSocio), hall.ExcluirControlAcceso,
				utils.NewNullString(hall.RegionCode), utils.IsHall(hall.IsHall), utils.IsExternal(hall.Externo))
			if err != nil {
				return err
			}
			hallsInserted++
			runtime.EventsEmit(ctx, "hallUpdated", fmt.Sprintf("Hall inserted: %s", hall.CodigoSalon))
		}
		progress := (float64(hallsInserted+hallsUpdated) / float64(len(halls.Body.ReadMultipleResult.ReadMultipleResult.CONFSALONES_WS))) * 100
		runtime.EventsEmit(ctx, "progress", progress)
	}

	if err != nil {
		runtime.EventsEmit(ctx, "hallUpdated", err.Error())
		return err
	}

	err = tx.Commit()
	if err != nil {
		runtime.EventsEmit(ctx, "hallUpdated", err.Error())
		return err
	}
	logger.Info("Data :", slog.Int("HallsInserted", hallsInserted), slog.Int("HallsUpdated", hallsUpdated))

	return nil
}
