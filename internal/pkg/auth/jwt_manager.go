package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type JWTManager struct {
	secretKey     string
	tokenDuration time.Duration
}

type UserClaims struct {
	jwt.StandardClaims
	UserName string `json:"username"`
	Role     Role   `json:"role"`
}

func NewJWTManager() *JWTManager {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Panic().Err(err).Msg("error random read")
	}
	secretKey := base64.StdEncoding.EncodeToString(randomBytes)
	log.Debug().Msg(secretKey)
	return &JWTManager{secretKey: secretKey,
		tokenDuration: 15 * time.Minute}
}

func (m *JWTManager) Generate(c *UserClaims) (string, error) {
	c.StandardClaims = jwt.StandardClaims{ExpiresAt: time.Now().Add(m.tokenDuration).Unix()}

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
