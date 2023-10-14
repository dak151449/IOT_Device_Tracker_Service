package authservice

import (
	"bytes"
	dao "iot-device-tracker-service/internal/app/dao/authservice"
	"iot-device-tracker-service/internal/pkg/auth"
	authapi "iot-device-tracker-service/pkg/api/auth_service"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/scrypt"
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

func isCorrectPassword(user *dao.User, password string) bool {
	checkPassword, err := scrypt.Key([]byte(password), []byte(user.Salt), 1<<15, 8, 1, 32)
	if err != nil {
		log.Debug().Err(err).Msg("error scrypt.Key check user password")
		return false
	}
	return bytes.Equal([]byte(user.HashedPassword), checkPassword)
}

func getClaims(user *dao.User) *auth.UserClaims {
	return &auth.UserClaims{
		UserName: user.UserName,
		Role:     auth.Role(user.Role),
	}
}
