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
func (i *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		log.Debug().Msgf("unary interceptor: %s", info.FullMethod)
		err := i.authorize(ctx, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(ctx, req)
	}
}

func (i *Interceptor) authorize(ctx context.Context, method string) error {
	accessibleRole, ok := i.accessibleRoles[method]
	if !ok {
		// everyone can access
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	claims, err := i.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	if accessibleRole >= claims.Role {
		return nil
	}

	return status.Error(codes.PermissionDenied, "no permission to access this RPC")
}
