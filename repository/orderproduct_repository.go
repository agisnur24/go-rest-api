package repository

import (
	"context"
	"database/sql"
	"go-rest-api/model/domain"
)

type OrderProductRepository interface {
	Save(ctx context.Context, tx *sql.Tx, odp domain.OrderProduct) domain.OrderProduct
	Update(ctx context.Context, tx *sql.Tx, odp domain.OrderProduct) domain.OrderProduct
	Delete(ctx context.Context, tx *sql.Tx, odp domain.OrderProduct)
	FindById(ctx context.Context, tx *sql.Tx, odpId int) (domain.OrderProduct, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct
}
