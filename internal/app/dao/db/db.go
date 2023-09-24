package dtservice_db

import (
	"iot-device-tracker-service/internal/db"
)

type DAO struct {
	db db.DB
}

func NewDAO(db db.DB) *DAO {
	return &DAO{
		db: db,
	}
}
