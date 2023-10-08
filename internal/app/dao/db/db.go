package appdb

import (
	"iot-device-tracker-service/internal/pkg/db"
)

type DAO struct {
	db db.DB
}

func NewDAO(db db.DB) *DAO {
	return &DAO{
		db: db,
	}
}
