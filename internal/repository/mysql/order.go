package models

import (
	"context"
	"database/sql"
	"log"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

type OrderRepository struct {
	DB *sql.DB
}

// NewOrderRepository will create an object that represent the order.Repository interface
func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (m *OrderRepository) getAll(ctx context.Context, query string, args ...interface{}) ([]domain.Order, int, error) {
	// Execute the SELECT query
	rows, err := m.DB.Query(query, args...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// Prepare a slice to hold the results
	var results []domain.Order
	var total int

	// Iterate through the rows
	for rows.Next() {
		var item domain.Order
		if err := rows.Scan(&item.Id,
			&item.ProductId,
			&item.Amount,
			&item.StatusId,
			&item.UserId,
			&item.Address,
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

func (m *OrderRepository) GetAllOrders(ctx context.Context, startDate, endDate, statusId, offset, limit string) ([]domain.Order, int, error) {
	query := "select * from orders as o " +
		"where 1=1 "

	if startDate != "" {
		startDate += " 00:00:00"
		query += " and o.created_at >= " + "'" + startDate + "'"
	}

	if endDate != "" {
		endDate += " 23:59:59"
		query += " and o.created_at <= " + "'" + endDate + "'"
	}

	if statusId != "" {
		query += " and o.status_id = " + statusId
	}

	query += " limit " + limit
	query += " offset " + offset
	return m.getAll(ctx, query)
}
