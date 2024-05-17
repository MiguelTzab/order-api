package repository

import (
	"context"
	"errors"

	"github.com/MiguelTzab/order-api/entity"
)

var ErrNotExist = errors.New("order does not exist")

type Repo interface {
	Insert(ctx context.Context, order entity.Order) error
	FindByID(ctx context.Context, id uint64) (entity.Order, error)
	DeleteByID(ctx context.Context, id uint64) error
	Update(ctx context.Context, order entity.Order) error
}
