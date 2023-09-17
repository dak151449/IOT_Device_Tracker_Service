package device_tracker_service

import (
	"google.golang.org/grpc"
	"proj/pkg/api"
)

type Implementation struct {
	api.UnimplementedDeviceTrackerServiceServer
}

func NewDeviceTrackerService() *Implementation {
	return &Implementation{}
}

func (i *Implementation) RegisterGRPC(s *grpc.Server) {
	api.RegisterDeviceTrackerServiceServer(s, i)
}
