package dtservice

import (
	"context"
	dao "iot-device-tracker-service/internal/app/dao/dtservice"
	"iot-device-tracker-service/internal/pkg/auth"
	dtapi "iot-device-tracker-service/pkg/api/device_tracker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) CreateDeviceGroup(ctx context.Context, req *dtapi.CreateDeviceGroupRequest) (*dtapi.CreateDeviceGroupResponse, error) {
	userID, err := auth.CheckUserRole(ctx, auth.User)
	if err != nil {
		return nil, err
	}

	if req.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid group name")
	}

	groups, err := i.dao.GetGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	groupNames := make(map[string]struct{}, len(groups))
	for _, v := range groups {
		groupNames[v.Name] = struct{}{}
	}

	if _, ok := groupNames[req.GetName()]; ok {
		return nil, status.Error(codes.AlreadyExists, "user already has a group with given name")
	}

	id, err := i.dao.CreateDeviceGroup(ctx, &dao.DeviceGroup{
		Name:        req.GetName(),
		UserID:      userID,
		Status:      req.GetStatus(),
		Description: req.GetDescription(),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &dtapi.CreateDeviceGroupResponse{
		Id: id,
	}, nil
}
