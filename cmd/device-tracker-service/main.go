package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	device_tracker_service "proj/internal/app/device-tracker-service"
	"proj/internal/pkg/app"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err).Msg("can't create app")
	}

	if err = a.Run(device_tracker_service.NewDeviceTrackerService()); err != nil {
		log.Fatal().Err(err).Msg("can't run app")
	}
}
