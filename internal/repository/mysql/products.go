package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

type ProductRepository struct {
	DB *sql.DB
}

// NewProductRepository will create an object that represent the product.Repository interface
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

func (m *ProductRepository) getAll(ctx context.Context, query string, args ...interface{}) ([]domain.Product, int, error) {
	// Execute the SELECT query
	rows, err := m.DB.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Prepare a slice to hold the results
	var results []domain.Product
	var total int

	// Iterate through the rows
	for rows.Next() {
		var item domain.Product
		if err := rows.Scan(&item.Id,
			&item.Style,
			&item.GenderId,
			&item.SizeId,
			&item.ColorId,
			&item.PatternId,
			&item.FigureId,
			&item.Price,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}
		results = append(results, item)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	total = len(results)
	return results, total, nil
}

func (m *ProductRepository) GetProductById(ctx context.Context, id int64) (domain.Product, error) {
	query := "select * from products where id = ? "
	return m.getOne(ctx, query, id)
}

func (m *ProductRepository) GetAllProducts(ctx context.Context, genderId, style, sizeId, offset, limit string) ([]domain.Product, int, error) {
	query := "select * from products as p " +
		"where 1=1 "

	if genderId != "" {
		query += "and p.gender_id  =  " + genderId
	}

	if style != "" {
		word := "'%" + style + "%'"
		query += " and p.`style` like " + word
	}

	if sizeId != "" {
		query += " and p.size_id  = " + sizeId
	}

	query += " limit " + limit
	query += " offset " + offset
	return m.getAll(ctx, query)

}
