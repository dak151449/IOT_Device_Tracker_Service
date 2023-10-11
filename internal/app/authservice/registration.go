package authservice

import (
	"context"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	auth_dao "iot-device-tracker-service/internal/app/dao/authservice"
	"iot-device-tracker-service/internal/pkg/auth"
	authapi "iot-device-tracker-service/pkg/api/auth_service"
)

func (server *AuthServer) Registration(ctx context.Context, req *authapi.RegistrationRequest) (*authapi.EmptyResponse, error) {
	user := req.GetUsername()
	password := req.GetPassword()
	hashedP, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	usr := auth_dao.User{UserName: user, HashedPassword: string(hashedP), Role: int(auth.User)} // какую роль ставить??
	//authserviceDao := authservice_db.NewDAO(db)
	_, err = server.dao.CreateUser(ctx, &usr)
	if err != nil {
		log.Warn().Err(err).Msgf("user already exists")
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	return &authapi.EmptyResponse{}, nil
}
