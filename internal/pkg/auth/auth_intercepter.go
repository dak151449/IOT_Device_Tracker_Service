package auth

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// Interceptor is a server interceptor for authentication and authorization
type Interceptor struct {
	jwtManager *JWTManager
}

// NewAuthInterceptor returns a new auth interceptor
func NewAuthInterceptor(jwtManager *JWTManager) *Interceptor {
	return &Interceptor{jwtManager}
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func (i *Interceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		userClaims, err := i.authorize(ctx)
		if err != nil {
			return nil, err
		}

		if userClaims != nil {
			log.Debug().Msgf("got user claims in auth interceptor: %s", info.FullMethod)
			ctx = context.WithValue(ctx, UserKey, userClaims)
		}

		return handler(ctx, req)
	}
}

// authorize checks if the user is authorized to access the method
func (i *Interceptor) authorize(ctx context.Context) (*UserClaims, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, nil
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, nil
	}

	accessToken := values[0]
	claims, err := i.jwtManager.Verify(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return claims, nil
}
