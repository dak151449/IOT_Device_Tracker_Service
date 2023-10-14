package auth

import (
	"context"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Role int

const (
	Admin Role = iota
	User
)

type contextKey string

const UserKey = contextKey("user")

type UserClaims struct {
	jwt.StandardClaims
	UserID   int64  `json:"id"`
	UserName string `json:"username"`
	Role     Role   `json:"role"`
}

// CheckUserRole returns status.Error
func CheckUserRole(ctx context.Context, acceptableRole Role) (int64, error) {
	value := ctx.Value(UserKey)
	if value == nil {
		return 0, status.Error(codes.Unauthenticated, "have no user claims in incoming context")
	}

	userClaims, ok := value.(*UserClaims)
	if !ok {
		return 0, status.Error(codes.Unauthenticated, "unable to extract user claims")
	}

	if userClaims.Role > acceptableRole {
		return 0, status.Error(codes.PermissionDenied, "no permission to access this RPC")
	}

	return userClaims.UserID, nil
}
