package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

const (
	DefaultSecretKey = "secretkey"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	UserID   int64  `json:"id"`
	UserName string `json:"username"`
	Role     Role   `json:"role"`
}

func NewJWTManager(secretKey string, tokenDuration time.Duration) *JWTManager {
	return &JWTManager{secretKey: secretKey,
		tokenDuration: tokenDuration}
}

func (m *JWTManager) Generate(c jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) Verify(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errors.New("unexpected token signing method")
			}
			return []byte(m.secretKey), nil
		})
	if err != nil {
		return nil, fmt.Errorf("invalid token %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
