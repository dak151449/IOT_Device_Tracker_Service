package authservice_db

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgconn"
	"github.com/pkg/errors"

	dao "iot-device-tracker-service/internal/app/dao/authservice"
	"iot-device-tracker-service/internal/pkg/db"
)

type DAO struct {
	db db.DB
}

var errUserAlreadyExists = errors.New("User with given username already exists")
var errDuplicateCode = "23505"

func NewDAO(db db.DB) *DAO {
	return &DAO{
		db: db,
	}
}

func (d *DAO) CreateUser(ctx context.Context, user *dao.User) (int64, error) {
	var id int64
	query, args, err := squirrel.Insert("users").
		Columns("username", "password", "salt", "email", "info", "role").
		Values(user.UserName, user.HashedPassword, user.Salt, user.Email, user.Info, user.Role).
		Suffix("RETURNING id").
		PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "failed to build SQL")
	}

	err = d.db.ExecQueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == errDuplicateCode {
				return 0, errUserAlreadyExists
			}
		}

		return 0, errors.Wrap(err, "db.ExecQueryRow error")
	}

	return id, nil
}

func (d *DAO) GetUser(ctx context.Context, username string) (*dao.User, error) {
	var user dao.User
	query, args, err := squirrel.Select("id", "username", "password", "salt", "email", "info", "role").
		From("users").
		Where(squirrel.Eq{"username": username}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "failed to build SQL")
	}

	err = d.db.Get(ctx, &user, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "db.Get error")
	}

	return &user, nil
}
