package authservice_dao

import "context"

type AuthServiceDAO interface {
	CreateUser(ctx context.Context, user *User) (int64, error)
	GetUser(ctx context.Context, username string) (*User, error)
}
