package dtservice

import (
	"context"
	"database/sql"
	"errors"
	dao "iot-device-tracker-service/internal/app/dao/dtservice"
	"iot-device-tracker-service/internal/pkg/auth"
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

func getUserIDFromContext(ctx context.Context) (int64, error) {
	userID, ok := ctx.Value(auth.UserIDKey).(int64)
	if !ok {
		return 0, errors.New("unable to extract user id from claims")
	}

	return userID, nil
}
