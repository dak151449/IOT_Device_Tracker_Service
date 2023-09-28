package dao

import "context"

type DTServiceDAO interface {
	Test(ctx context.Context) error
}
