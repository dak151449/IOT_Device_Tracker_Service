package dtservice

import (
	"context"
	dtapi "iot-device-tracker-service/pkg/api/device_tracker"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetDevicesFromGroup(ctx context.Context, req *dtapi.GetDevicesFromGroupRequest) (*dtapi.GetDevicesFromGroupResponse, error) {
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
