package dtservice

import (
	"iot-device-tracker-service/internal/app/dao"
	"iot-device-tracker-service/pkg/api"

	"google.golang.org/grpc"
)

type Implementation struct {
	api.UnimplementedDeviceTrackerServiceServer

	dao dao.DTServiceDAO
}

func NewDeviceTrackerService(dao dao.DTServiceDAO) *Implementation {
	return &Implementation{dao: dao}
}

func (i *Implementation) RegisterGRPC(s *grpc.Server) {
	api.RegisterDeviceTrackerServiceServer(s, i)
}
