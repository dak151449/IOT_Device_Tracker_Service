package dtservice_db

import (
	"context"
	dao "iot-device-tracker-service/internal/app/dao/dtservice"
	"iot-device-tracker-service/internal/pkg/db"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

type DAO struct {
	db db.DB
}

func NewDAO(db db.DB) *DAO {
	return &DAO{
		db: db,
	}
}

func (d *DAO) GetGroupsByUserID(ctx context.Context, userID int64) ([]dao.DeviceGroup, error) {
	var groups []dao.DeviceGroup

	query, args, err := squirrel.Select("id", "name", "status", "created_at", "description").
		From("device_groups").
		Where(squirrel.Eq{"user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build SQL")
	}

	err = d.db.Select(ctx, &groups, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "db.Select error")
	}

	return groups, nil
}

func (d *DAO) GetDevicesByGroupID(ctx context.Context, userID int64) ([]dao.Device, error) {
	var devices []dao.Device

	query, args, err := squirrel.Select("id", "name", "status", "created_at", "description").
		From("devices").
		Where(squirrel.Eq{"device_group_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build SQL")
	}

	err = d.db.Select(ctx, &devices, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "db.Select error")
	}

	return devices, nil
}
