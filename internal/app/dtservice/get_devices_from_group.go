package dtservice

import (
	"context"
	"iot-device-tracker-service/internal/pkg/auth"
	dtapi "iot-device-tracker-service/pkg/api/device_tracker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetDevicesFromGroup(ctx context.Context, req *dtapi.GetDevicesFromGroupRequest) (*dtapi.GetDevicesFromGroupResponse, error) {
	userID, err := auth.CheckUserRole(ctx, auth.User)
	if err != nil {
		return nil, err
	}

	if err = i.checkGroupExistence(ctx, userID, req.GetGroupId()); err != nil {
		return nil, err
	}

	devices, err := i.dao.GetDevicesByGroupID(ctx, req.GetGroupId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]*dtapi.DeviceData, 0, len(devices))
	for _, g := range devices {
		res = append(res, &dtapi.DeviceData{
			Id:          g.ID,
			Name:        g.Name,
			Status:      g.Status,
			CreatedAt:   toTimestampPB(g.CreatedAt),
			Description: g.Description,
		})
	}

	return &dtapi.GetDevicesFromGroupResponse{
		Devices: res,
	}, nil
}

// checkGroupExistence returns status.Error
func (i *Implementation) checkGroupExistence(ctx context.Context, userID int64, groupID int64) error {
	groups, err := i.dao.GetGroupsByUserID(ctx, userID)
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}

	groupIDs := make(map[int64]struct{}, len(groups))
	for _, v := range groups {
		groupIDs[v.ID] = struct{}{}
	}

	if _, ok := groupIDs[groupID]; !ok {
		return status.Error(codes.NotFound, "device group not found for user")
	}

	return nil
}
