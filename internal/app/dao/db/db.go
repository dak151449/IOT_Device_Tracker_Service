package dtservice_db

import (
	"context"
	"fmt"
	"iot-device-tracker-service/internal/app/dao"
	"iot-device-tracker-service/internal/db"

	"github.com/Masterminds/squirrel"
	"github.com/rs/zerolog/log"
)

const (
	usersTable = "users"
)

type DAO struct {
	db db.DB
}

func NewDAO(db db.DB) *DAO {
	return &DAO{
		db: db,
	}
}

func (d *DAO) Test(ctx context.Context) error {
	var u dao.User

	query, args, err := squirrel.Select("id", "name", "created_at", "updated_at").
		From(usersTable).
		Where(squirrel.Eq{"id": 123}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Info().Err(err).Msg("failed to build SQL")
		return err
	}

	err = d.db.Get(ctx, &u, query, args...)
	if err != nil {
		return err
	}

	fmt.Println(u)
	return nil
}
