package mysql

import (
	"context"
	"database/sql"
	"log"
	"time"

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
		if err := rows.Scan(
			&item.Id,
			&item.ProductId,
			&item.Amount,
			&item.UserId,
			&item.Address,
			&item.StatusId,
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
	query := "select o.id,o.product_id,o.amount,o.user_id,o.address,h.status_id,o.created_at,o.updated_at from orders as o " +
		" left join history as h on h.order_id = o.id " +
		" where 1=1 "

	if statusId != "" {
		query += " and h.status_id = " + statusId
	} else {
		query += " and h.status_id in (1,2,3,4)"
	}

	if startDate != "" {
		startDate += " 00:00:00"
		query += " and o.created_at >= " + "'" + startDate + "'"
	}

	if endDate != "" {
		endDate += " 23:59:59"
		query += " and o.created_at <= " + "'" + endDate + "'"
	}

	query += " limit " + limit
	query += " offset " + offset
	return m.getAll(ctx, query)
}

func (m *OrderRepository) InsertOrder(ctx context.Context, o *domain.Order) (err error) {
	query := "insert orders set product_id=? ,amount=?,user_id=?,address=?,updated_at=?,created_at=? "
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, o.ProductId, o.Amount, o.UserId, o.Address, o.UpdatedAt, o.CreatedAt)
	if err != nil {
		return
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return
	}
	o.Id = int(lastId)

	// table history
	query2 := "insert history set order_id=?,status_id=?,created_at=?,updated_at=?"
	stmt2, err := m.DB.PrepareContext(ctx, query2)
	if err != nil {
		return
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt2.ExecContext(ctx, o.Id, 1, now, now)
	if err != nil {
		return
	}
	return
}
