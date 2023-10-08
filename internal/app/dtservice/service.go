package dtservice

import (
	"database/sql"
	"iot-device-tracker-service/internal/app/dao"
	dtapi "iot-device-tracker-service/pkg/api/device_tracker"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Implementation struct {
	dtapi.UnimplementedDeviceTrackerServiceServer

	dao dao.DTServiceDAO
}

func NewDeviceTrackerService(dao dao.DTServiceDAO) *Implementation {
	return &Implementation{dao: dao}
}

func (i *Implementation) RegisterGRPC(s *grpc.Server) {
	dtapi.RegisterDeviceTrackerServiceServer(s, i)
}

func toTimestampPB(t sql.NullTime) *timestamppb.Timestamp {
	if t.Valid {
		return timestamppb.New(t.Time)
	}
	return nil
}
