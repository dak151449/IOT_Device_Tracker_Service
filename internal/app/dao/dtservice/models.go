package dtservice_dao

import "database/sql"

type DeviceGroup struct {
	ID          int64        `db:"id"`
	Name        string       `db:"name"`
	UserID      int64        `db:"user_id"`
	Status      string       `db:"status"`
	CreatedAt   sql.NullTime `db:"created_at"`
	Description string       `db:"description"`
}

type Device struct {
	ID          int64        `db:"id"`
	Name        string       `db:"name"`
	GroupID     int64        `db:"device_group_id"`
	Status      string       `db:"status"`
	CreatedAt   sql.NullTime `db:"created_at"`
	Description string       `db:"description"`
}
