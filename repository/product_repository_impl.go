package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"go-rest-api/helper"
	"go-rest-api/model/domain"
	"go-rest-api/model/web"
)

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (p ProductRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	logrus.Info("Repo Save Start")
	SQL := "insert into product(name, price, category_id) values (?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	product.Id = int(id)
	logrus.Info("Repo Save End")
	return product
}

func (p ProductRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	logrus.Info("Repo Update Start")
	SQL := "update product set name = ?, price = ?, category_id = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Name, product.Price, product.CategoryId, product.Id)
	helper.PanicIfError(err)

	logrus.Info("Repo Update End")
	return product
}

func (p ProductRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, product domain.Product) {
	logrus.Info("Repo Find by Id Start")
	SQL := "delete from product where id = ?"
	_, err := tx.ExecContext(ctx, SQL, product.Id)
	helper.PanicIfError(err)
	logrus.Info("Repo Find by Id End")
}

func (p ProductRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (web.ProductResponse, error) {
	logrus.Info("Repo Find by Id Start")
	//SELECT a.*, b.name as 'category_name' FROM `product` a
	//INNER JOIN category b on a.category_id=b.id;
	SQL := "select p.id, p.name, p.price, p.category_id, c.name from product p inner join category c on p.category_id=c.id where p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	defer rows.Close()

	product := web.ProductResponse{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		logrus.Info("Repo Find by Id End")
		return product, nil
	} else {
		logrus.Info("Repo Find by Id End")
		return product, errors.New("Product is not found")
	}
}

func (p ProductRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []web.ProductResponse {
	logrus.Info("Repo Find All Start")
	//SELECT a.*, b.name as 'category_name' FROM `product` a
	//INNER JOIN category b on a.category_id=b.id;
	SQL := "select p.id, p.name, p.price, p.category_id, c.name from product p inner join category c on p.category_id=c.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var products []web.ProductResponse
	for rows.Next() {
		product := web.ProductResponse{}
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.CategoryId, &product.CategoryName)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	logrus.Info("Repo Find All End")
	return products
}
