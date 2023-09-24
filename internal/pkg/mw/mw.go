package mw

import (
	"context"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func LogResponseInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	resp, err := handler(ctx, req)

	if err != nil {
		log.Info().Err(err).Msgf("gRPC error in method %s", info.FullMethod)
	} else {
		log.Debug().Msgf("gRPC method %s executed successfully", info.FullMethod)
	}

	return resp, err
}
