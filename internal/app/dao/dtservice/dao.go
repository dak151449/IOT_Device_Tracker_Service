package dtservice_dao

import "context"

type DTServiceDAO interface {
	GetGroupsByUserID(ctx context.Context, userID int64) ([]DeviceGroup, error)
	GetDevicesByGroupID(ctx context.Context, userID int64) ([]Device, error)
	CreateDeviceGroup(ctx context.Context, group *DeviceGroup) (int64, error)
}
