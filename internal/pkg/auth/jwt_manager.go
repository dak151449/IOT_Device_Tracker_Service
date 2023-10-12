package auth

import (
	"crypto/rand"
	"encoding/base64"
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
	UserID   int64  `json:"id"`
	UserName string `json:"username"`
	Role     Role   `json:"role"`
}

func NewJWTManager(tokenDuration time.Duration) *JWTManager {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Panic().Err(err).Msg("error random read")
	}

	secretKey := base64.StdEncoding.EncodeToString(randomBytes)
	return &JWTManager{secretKey: secretKey,
		tokenDuration: tokenDuration}
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
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, errors.New("unable to extract user claims")
	}

	return claims, nil
}
