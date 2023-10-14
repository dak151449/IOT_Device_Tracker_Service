package authservice

import (
	"bytes"
	"context"
	"errors"
	dao "iot-device-tracker-service/internal/app/dao/authservice"
	"iot-device-tracker-service/internal/pkg/auth"
	authapi "iot-device-tracker-service/pkg/api/auth_service"

	"github.com/jackc/pgx/v4"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/scrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Login(ctx context.Context, req *authapi.LoginRequest) (*authapi.LoginResponse, error) {
	user, err := i.dao.GetUser(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "user does not exist")
		}

		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !isCorrectPassword(user, req.GetPassword()) {
		return nil, status.Error(codes.NotFound, "invalid password")
	}

	token, err := i.jwtManager.Generate(getClaims(user))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate access token: %v", err)
	}

	res := &authapi.LoginResponse{AccessToken: token}
	return res, nil
}

func isCorrectPassword(user *dao.User, password string) bool {
	checkPassword, err := scrypt.Key([]byte(password), user.Salt, scryptN, scryptR, scryptP, scryptKeyLen)
	if err != nil {
		log.Debug().Err(err).Msg("scrypt.Key error checking user password")
		return false
	}

	return bytes.Equal(user.HashedPassword, checkPassword)
}

func getClaims(user *dao.User) *auth.UserClaims {
	return &auth.UserClaims{
		UserID:   user.ID,
		UserName: user.UserName,
		Role:     auth.Role(user.Role),
	}
}
