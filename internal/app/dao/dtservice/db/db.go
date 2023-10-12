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

func (d *DAO) CreateDeviceGroup(ctx context.Context, group *dao.DeviceGroup) (int64, error) {
	var id int64
	query, args, err := squirrel.Insert("device_groups").
		Columns("name", "user_id", "status", "description").
		Values(group.Name, group.UserID, group.Status, group.Description).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build SQL")
	}

	err = d.db.ExecQueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return 0, errors.Wrap(err, "db.ExecQueryRow error")
	}

	return id, nil
}
