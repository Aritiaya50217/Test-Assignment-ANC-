package mysql

import (
	"context"
	"database/sql"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

type ColorRepository struct {
	DB *sql.DB
}

// NewColorRepository will create an object that represent the color.Repository interface
func NewColorRepository(db *sql.DB) *ColorRepository {
	return &ColorRepository{
		DB: db,
	}
}

func (m *ColorRepository) getOne(ctx context.Context, query string, args ...interface{}) (res domain.Color, err error) {
	stmt, err := m.DB.PrepareContext(ctx, query)
	if err != nil {
		return domain.Color{}, err
	}
	row := stmt.QueryRowContext(ctx, args...)
	res = domain.Color{}

	err = row.Scan(
		&res.Id,
		&res.Name,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	return
}

func (m *ColorRepository) GetColorById(ctx context.Context, id int64) (domain.Color, error) {
	query := `SELECT * FROM colors WHERE id=?`
	return m.getOne(ctx, query, id)
}
