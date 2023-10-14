package dtservice

import (
	"context"
	dao "iot-device-tracker-service/internal/app/dao/dtservice"
	dtapi "iot-device-tracker-service/pkg/api/device_tracker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateDevice(ctx context.Context, req *dtapi.CreateDeviceRequest) (*dtapi.CreateDeviceResponse, error) {
	userID, err := getUserIDFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid device name")
	}

	groups, err := i.dao.GetGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	groupIDs := make(map[int64]struct{}, len(groups))
	for _, v := range groups {
		groupIDs[v.ID] = struct{}{}
	}

	if _, ok := groupIDs[req.GetGroupId()]; !ok {
		return nil, status.Error(codes.NotFound, "device group not found")
	}

	devices, err := i.dao.GetDevicesByGroupID(ctx, req.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	deviceNames := make(map[string]struct{}, len(devices))
	for _, v := range devices {
		deviceNames[v.Name] = struct{}{}
	}

	if _, ok := deviceNames[req.GetName()]; ok {
		return nil, status.Error(codes.AlreadyExists, "group already has a device with given name")
	}

	id, err := i.dao.CreateDevice(ctx, &dao.Device{
		Name:        req.GetName(),
		GroupID:     req.GetGroupId(),
		Status:      req.GetStatus(),
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &dtapi.CreateDeviceResponse{
		Id: id,
	}, nil
}
