package authservice

import (
	"context"
	auth_dao "iot-device-tracker-service/internal/app/dao/authservice"
	"iot-device-tracker-service/internal/pkg/auth"
	authapi "iot-device-tracker-service/pkg/api/auth_service"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Registration(ctx context.Context, req *authapi.RegistrationRequest) (*authapi.EmptyResponse, error) {
	username := req.GetUsername()
	password := req.GetPassword()
	hashedP, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := auth_dao.User{UserName: username, HashedPassword: string(hashedP), Role: int(auth.User)}

	_, err = i.dao.CreateUser(ctx, &user)
	if err != nil {
		return nil, status.Error(codes.AlreadyExists, err.Error())
	}

	return &authapi.EmptyResponse{}, nil
}
