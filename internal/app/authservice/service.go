package authservice

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"

	dao "iot-device-tracker-service/internal/app/dao/authservice"
	"iot-device-tracker-service/internal/pkg/auth"
	authapi "iot-device-tracker-service/pkg/api/auth_service"
)

type AuthServer struct {
	authapi.UnimplementedAuthServiceServer

	dao        dao.AuthServiceDAO
	jwtManager *auth.JWTManager
}

func NewAuthServer(userStore dao.AuthServiceDAO, jwtManager *auth.JWTManager) *AuthServer {
	return &AuthServer{
		dao:        userStore,
		jwtManager: jwtManager,
	}
}

func (server *AuthServer) RegisterGRPC(i *grpc.Server) {
	authapi.RegisterAuthServiceServer(i, server)
}

func isCorrectPassword(user *dao.User, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

func getClaims(user *dao.User) jwt.Claims {
	return &auth.UserClaims{
		UserName: user.UserName,
		Role:     auth.Role(user.Role),
	}
}
