package main

import (
	"context"
	dtservice_db "iot-device-tracker-service/internal/app/dao/db"
	"iot-device-tracker-service/internal/app/dtservice"
	"iot-device-tracker-service/internal/pkg/app"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	ctx := context.Background()

	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err).Msg("can't create app")
	}

	db, err := dtservice_db.Connect(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("can't connect to postgres")
	}
	defer db.GetPool(ctx).Close()

	dao := dtservice_db.NewDAO(db)

	err = dao.Test(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("dao err")
	}

	if err = a.Run(dtservice.NewDeviceTrackerService(dao)); err != nil {
		log.Fatal().Err(err).Msg("can't run app")
	}
}
