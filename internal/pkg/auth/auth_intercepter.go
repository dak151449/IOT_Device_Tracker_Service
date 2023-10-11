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
func (interceptor *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Info().Msgf("--> unary interceptor: %s", info.FullMethod)
		err := interceptor.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (interceptor *Interceptor) authorize(ctx context.Context, method string) error {
	accessibleRole, ok := interceptor.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Warn().Msgf("metadata is not provided")
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		log.Warn().Msgf("authorization token is not provided")
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		log.Warn().Err(err).Msgf("access token is invalid: %v", err)
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if accessibleRole >= claims.Role {
		return nil
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
