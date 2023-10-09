package dtservice

import (
	"context"
	dtapi "iot-device-tracker-service/pkg/api/device_tracker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetDeviceGroups(ctx context.Context, req *dtapi.GetDeviceGroupsRequest) (*dtapi.GetDeviceGroupsResponse, error) {
	groups, err := i.dao.GetGroupsByUserID(ctx, req.GetUserId())
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
