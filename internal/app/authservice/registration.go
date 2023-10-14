package authservice

import (
	"context"
	"crypto/rand"
	auth_dao "iot-device-tracker-service/internal/app/dao/authservice"
	auth_db "iot-device-tracker-service/internal/app/dao/authservice/db"
	"iot-device-tracker-service/internal/pkg/auth"
	authapi "iot-device-tracker-service/pkg/api/auth_service"
	"unicode/utf8"

	"github.com/pkg/errors"
	"golang.org/x/crypto/scrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	minLenLogin    = 4
	minLenPassword = 4
)

const (
	scryptN      = 1 << 15
	scryptR      = 8
	scryptP      = 1
	scryptKeyLen = 32
	saltSize     = 32
)

func (i *Implementation) Registration(ctx context.Context, req *authapi.RegistrationRequest) (*authapi.EmptyResponse, error) {
	if utf8.RuneCountInString(req.GetUsername()) < minLenLogin {
		return nil, status.Errorf(codes.InvalidArgument, "login must be at least %d characters", minLenLogin)
	}

	if utf8.RuneCountInString(req.GetPassword()) < minLenPassword {
		return nil, status.Errorf(codes.InvalidArgument, "password must be at least %d characters", minLenPassword)
	}

	randomBytes, hashedP, err := generateHash(req.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	user := auth_dao.User{UserName: req.GetUsername(), HashedPassword: hashedP, Role: int(auth.User), Salt: randomBytes}

	_, err = i.dao.CreateUser(ctx, &user)
	if err != nil {
		if errors.Is(err, auth_db.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &authapi.EmptyResponse{}, nil
}

func generateHash(password string) ([]byte, []byte, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error generating salt")
	}

	hashedPassword, err := scrypt.Key([]byte(password), salt, scryptN, scryptR, scryptP, scryptKeyLen)
	if err != nil {
		return nil, nil, errors.Wrap(err, "error generating password hash")
	}

	return salt, hashedPassword, nil
}
