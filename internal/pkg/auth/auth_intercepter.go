package auth

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Role int

type contextKey string

const UserIDKey = contextKey("userID")

const (
	Admin Role = iota
	User
)

// Interceptor is a server interceptor for authentication and authorization
type Interceptor struct {
	jwtManager      *JWTManager
	accessibleRoles map[string]Role
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager *JWTManager, accessibleRoles map[string]Role) *Interceptor {
	return &Interceptor{jwtManager, accessibleRoles}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (i *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Info().Msgf("--> unary interceptor: %s", info.FullMethod)
		userID, err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, UserIDKey, userID)

		return handler(ctx, req)
	}
}

func (i *Interceptor) authorize(ctx context.Context, method string) (int64, error) {
	accessibleRole, ok := i.accessibleRoles[method]
	if !ok {
		// everyone can access
		return 0, nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return 0, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return 0, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := i.jwtManager.Verify(accessToken)
	if err != nil {
		return 0, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if accessibleRole >= claims.Role {
		return claims.UserID, nil
	}

	return 0, status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
