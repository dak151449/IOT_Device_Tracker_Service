package authservice

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	authapi "iot-device-tracker-service/pkg/api/auth_service"
)

func (server *AuthServer) Login(ctx context.Context, req *authapi.LoginRequest) (*authapi.LoginResponse, error) {
	user, err := server.dao.GetUser(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	if user == nil || !isCorrectPassword(user, req.GetPassword()) {
		log.Warn().Err(err).Msgf("incorrect username/password")
		return nil, status.Error(codes.NotFound, "incorrect username/password")
	}
	token, err := server.jwtManager.Generate(getClaims(user))
	if err != nil {
		log.Warn().Err(err).Msgf("cannot find user")
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}
	res := &authapi.LoginResponse{AccessToken: token}
	return res, nil
}
