package mysql

import (
	"context"
	"database/sql"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

type HistoryRepository struct {
	DB *sql.DB
}

func NewHistoryRepository(db *sql.DB) *HistoryRepository {
	return &HistoryRepository{
		DB: db,
	}
}

func (m *HistoryRepository) InsertHistory(ctx context.Context, h *domain.History) (err error) {
	query := "insert history set order_id,status_id ,created_at,updated_at"
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, h.OrderId, h.StatusId, h.CreatedAt, h.UpdatedAt)
	if err != nil {
		return
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return
	}
	h.Id = int(lastId)
	return
}
