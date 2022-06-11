package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
)

type OrderProductRepositoryImpl struct {
}

func NewOrderProductRepository() OrderProductRepository {
	return &OrderProductRepositoryImpl{}
}

func (op OrderProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, odp domain.OrderProduct) domain.OrderProduct {
	SQL := "insert into order_product(order_id, product_id, qty, price, amount) values (?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, odp.OrderId, odp.ProductId, odp.Qty, odp.Price, odp.Amount)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	odp.Id = int(id)
	return odp
}

func (op OrderProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, odp domain.OrderProduct) domain.OrderProduct {
	SQL := "update order_product set order_id = ?, product_id = ?, qty = ?, price = ?, amount = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, odp.OrderId, odp.ProductId, odp.Qty, odp.Price, odp.Amount, odp.Id)
	helper.PanicIfError(err)

	return odp
}

func (op OrderProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, odp domain.OrderProduct) {
	SQL := "delete from order_product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, odp.Id)
	helper.PanicIfError(err)
}

func (op OrderProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, odpId int) (domain.OrderProduct, error) {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, odpId)
	helper.PanicIfError(err)
	defer rows.Close()

	odp := domain.OrderProduct{}
	if rows.Next() {
		err := rows.Scan(&odp.Id, &odp.OrderId, &odp.ProductId, &odp.Qty, &odp.Price, &odp.Amount)
		helper.PanicIfError(err)
		return odp, nil
	} else {
		return odp, errors.New("Order-Product is not found")
	}
}

func (op OrderProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.OrderProduct {
	SQL := "select id, order_id, product_id, qty, price, amount from order_product"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderproducts []domain.OrderProduct
	for rows.Next() {
		odp := domain.OrderProduct{}
		err := rows.Scan(&odp.Id, &odp.OrderId, &odp.ProductId, &odp.Qty, &odp.Price, &odp.Amount)
		helper.PanicIfError(err)
		orderproducts = append(orderproducts, odp)
	}
	return orderproducts
}
