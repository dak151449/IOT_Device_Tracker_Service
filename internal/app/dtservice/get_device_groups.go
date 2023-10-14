package dtservice

import (
	"context"
	"iot-device-tracker-service/internal/pkg/auth"
	dtapi "iot-device-tracker-service/pkg/api/device_tracker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetDeviceGroups(ctx context.Context, _ *dtapi.GetDeviceGroupsRequest) (*dtapi.GetDeviceGroupsResponse, error) {
	userID, err := auth.CheckUserRole(ctx, auth.User)
	if err != nil {
		return nil, err
	}

	groups, err := i.dao.GetGroupsByUserID(ctx, userID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	res := make([]*dtapi.DeviceGroupData, 0, len(groups))
	for _, g := range groups {
		res = append(res, &dtapi.DeviceGroupData{
			Id:          g.ID,
			Name:        g.Name,
			Status:      g.Status,
			CreatedAt:   toTimestampPB(g.CreatedAt),
			Description: g.Description,
		})
	}

	return &dtapi.GetDeviceGroupsResponse{
		Groups: res,
	}, nil
}
