package main

import (
	"context"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"

	"iot-device-tracker-service/internal/app/authservice"
	authservice_db "iot-device-tracker-service/internal/app/dao/authservice/db"
	dtservice_db "iot-device-tracker-service/internal/app/dao/dtservice/db"
	"iot-device-tracker-service/internal/app/dtservice"
	"iot-device-tracker-service/internal/config"
	"iot-device-tracker-service/internal/pkg/app"
	"iot-device-tracker-service/internal/pkg/auth"
	"iot-device-tracker-service/internal/pkg/db"
)

func accessibleRoles() map[string]auth.Role {
	const laptopServicePath = "/device_tracker.DeviceTrackerService/"

	return map[string]auth.Role{
		laptopServicePath + "GetDeviceGroups":     auth.User,
		laptopServicePath + "GetDevicesFromGroup": auth.User,
	}
}

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
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

	dtserviceDao := dtservice_db.NewDAO(db)
	authserviceDao := authservice_db.NewDAO(db)

	tokenDuration, err := config.GetTokenDuration()
	if err != nil {
		log.Fatal().Err(err).Msg("can't parse tokenDuration")
	}

	jwtManager := auth.NewJWTManager(auth.SecretKey, tokenDuration)
	AuthServer := authservice.NewAuthServer(authserviceDao, jwtManager)
	authIntercepter := auth.NewAuthInterceptor(jwtManager, accessibleRoles())

	options := []grpc.UnaryServerInterceptor{authIntercepter.Unary()}

	if err = a.Run(options, dtservice.NewDeviceTrackerService(dtserviceDao), AuthServer); err != nil {
		log.Fatal().Err(err).Msg("can't run app")
	}
}
