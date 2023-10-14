package authservice

import (
	dao "iot-device-tracker-service/internal/app/dao/authservice"
	"iot-device-tracker-service/internal/pkg/auth"
	authapi "iot-device-tracker-service/pkg/api/auth_service"

	"google.golang.org/grpc"
)

type Implementation struct {
	authapi.UnimplementedAuthServiceServer

	dao        dao.AuthServiceDAO
	jwtManager *auth.JWTManager
}

func NewAuthService(userStore dao.AuthServiceDAO, jwtManager *auth.JWTManager) *Implementation {
	return &Implementation{
		dao:        userStore,
		jwtManager: jwtManager,
	}
}

func (i *Implementation) RegisterGRPC(s *grpc.Server) {
	authapi.RegisterAuthServiceServer(s, i)
}
