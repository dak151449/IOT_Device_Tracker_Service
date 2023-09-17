package device_tracker_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"proj/pkg/api"
	"strings"
)

func (i *Implementation) ToUpper(ctx context.Context, req *api.ToUpperRequest) (*api.ToUpperResponse, error) {
	if req.GetInputString() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty input string")
	}

	return &api.ToUpperResponse{
		ConvertedString: strings.ToUpper(req.GetInputString()),
	}, nil
}
