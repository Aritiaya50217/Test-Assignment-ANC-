package mysql_test

import (
	"context"
	"testing"
	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	productMysqlRepo "github.com/Aritiaya50217/Test-Assignment-ANC/internal/repository/mysql"
	"github.com/stretchr/testify/assert"
)

func TestGetProductById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	create := time.Now().Format("2006-01-02 15:04:05")
	update := time.Now().Format("2006-01-02 15:04:05")

	rows := sqlmock.NewRows([]string{"id", "style", "gender_id", "size_id", "color_id", "pattern_id", "figure_id", "price", "updated_at", "created_at"}).
		AddRow(1, "T-Shirt", 1, 1, 1, 1, 1, 400, create, update)

	query := "select * from products where id = ?"
	mock.ExpectQuery(query).WillReturnRows(rows)
	p := productMysqlRepo.NewProductRepository(db)

	num := int64(5)
	product, _ := p.GetProductById(context.TODO(), num)
	assert.NotNil(t, product)
}
