package models

import (
	"context"
	"database/sql"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

type ProductRepository struct {
	DB *sql.DB
}

// NewProductRepository will create an object that represent the product.Repository inteface
func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{
		DB: db,
	}
}

func (m *ProductRepository) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Product, err error) {
	smtp, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Product{}, err
	}
	row := smtp.QueryRowContext(ctx, args...)
	res = domain.Product{}

	err = row.Scan(
		&res.Id,
		&res.Style,
		&res.GenderId,
		&res.SizeId,
		&res.ColorId,
		&res.PatternId,
		&res.FigureId,
		&res.Price,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return
}

func (m *ProductRepository) GetProductById(ctx context.Context, id int64) (domain.Product, error) {
	query := "select * from products where id = ? "
	return m.getOne(ctx, query, id)
}
