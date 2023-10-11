package auth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

const (
	SecretKey = "secretkey"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey,
		tokenDuration: tokenDuration}
}

func (manager *JWTManager) Generate(c jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(manager.secretKey))
}

type UserClaims struct {
	jwt.StandardClaims
	UserName string `json:"username"`
	Role     Role   `json:"role"`
}

func (manager *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				log.Warn().Msgf("unexpected token signing method")
				return nil, errors.New("unexpected token signing method")
			}
			return []byte(manager.secretKey), nil
		})
	if err != nil {
		log.Warn().Err(err).Msgf("invalid token")
		return nil, errors.New("invalid token	" + err.Error())
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		log.Warn().Msgf("invalid token")
		return nil, errors.New("invalid token	")
	}
	return claims, nil
}
