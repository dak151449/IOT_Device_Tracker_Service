package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"iot-device-tracker-service/internal/app/dt_service"
	"iot-device-tracker-service/internal/pkg/app"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err).Msg("can't create app")
	}

	if err = a.Run(dt_service.NewDeviceTrackerService()); err != nil {
		log.Fatal().Err(err).Msg("can't run app")
	}
}
