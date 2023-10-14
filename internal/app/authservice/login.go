package authservice

import (
	"context"
	authapi "iot-device-tracker-service/pkg/api/auth_service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *authapi.LoginRequest) (*authapi.LoginResponse, error) {
	user, err := i.dao.GetUser(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !isCorrectPassword(user, req.GetPassword()) {
		return nil, status.Error(codes.NotFound, "incorrect username/password")
	}

	token, err := i.jwtManager.Generate(getClaims(user))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	res := &authapi.LoginResponse{AccessToken: token}
	return res, nil
}
