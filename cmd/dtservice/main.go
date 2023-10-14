package main

import (
	"context"
	"iot-device-tracker-service/internal/app/authservice"
	auth_db "iot-device-tracker-service/internal/app/dao/authservice/db"
	dt_db "iot-device-tracker-service/internal/app/dao/dtservice/db"
	"iot-device-tracker-service/internal/app/dtservice"
	"iot-device-tracker-service/internal/pkg/app"
	"iot-device-tracker-service/internal/pkg/auth"
	"iot-device-tracker-service/internal/pkg/config"
	"iot-device-tracker-service/internal/pkg/db"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	ctx := context.Background()

	a, err := app.New()
	if err != nil {
		log.Fatal().Err(err).Msg("can't create app")
	}

	db, err := db.Connect(ctx)
	if err != nil {
		log.Fatal().Err(err).Msg("can't connect to postgres")
	}
	defer db.GetPool(ctx).Close()

	dtDAO := dt_db.NewDAO(db)
	authDAO := auth_db.NewDAO(db)

	jwtTokenDuration, err := config.GetJWTTokenDuration()
	if err != nil {
		log.Fatal().Err(err).Msg("can't parse jwtTokenDuration")
	}

	jwtManager := auth.NewJWTManager(jwtTokenDuration)
	authInterceptor := auth.NewAuthInterceptor(jwtManager)

	if err = a.Run(
		[]grpc.UnaryServerInterceptor{authInterceptor.Unary()},
		dtservice.NewDeviceTrackerService(dtDAO),
		authservice.NewAuthService(authDAO, jwtManager)); err != nil {
		log.Fatal().Err(err).Msg("can't run app")
	}
}
