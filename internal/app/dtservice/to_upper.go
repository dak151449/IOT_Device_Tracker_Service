package dtservice

import (
	"context"
	"iot-device-tracker-service/pkg/api"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) ToUpper(_ context.Context, req *api.ToUpperRequest) (*api.ToUpperResponse, error) {
	if req.GetInputString() == "" {
		return nil, status.Error(codes.InvalidArgument, "empty input string")
	}

	return &api.ToUpperResponse{
		ConvertedString: strings.ToUpper(req.GetInputString()),
	}, nil
}
