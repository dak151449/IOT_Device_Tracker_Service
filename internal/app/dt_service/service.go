package dt_service

import (
	"google.golang.org/grpc"
	"iot-device-tracker-service/pkg/api"
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
