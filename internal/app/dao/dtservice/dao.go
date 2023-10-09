package dtservice_dao

import "context"

type DTServiceDAO interface {
	GetGroupsByUserID(ctx context.Context, userID int64) ([]DeviceGroup, error)
	GetDevicesByGroupID(ctx context.Context, userID int64) ([]Device, error)
}
