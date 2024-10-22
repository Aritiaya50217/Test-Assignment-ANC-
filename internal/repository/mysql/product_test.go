package mysql_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
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

func TestGetProducts(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	create := time.Now().Format("2006-01-02 15:04:05")
	update := time.Now().Format("2006-01-02 15:04:05")

	mockProducts := []domain.Product{
		{
			Id: 1, Style: "T-Shirt", GenderId: 1, SizeId: 1, ColorId: 1, Price: "400", PatternId: 1,
			FigureId: 1, UpdatedAt: update, CreatedAt: create,
		},
		{
			Id: 2, Style: "T-Shirt", GenderId: 1, SizeId: 2, ColorId: 1, Price: "400", PatternId: 2,
			FigureId: 2, UpdatedAt: update, CreatedAt: create,
		},
	}
	priceZero, _ := strconv.Atoi(mockProducts[0].Price)
	priceOne, _ := strconv.Atoi(mockProducts[1].Price)

	rows := sqlmock.NewRows([]string{"id", "style", "gender_id", "size_id", "color_id",
		"price", "pattern_id", "figure_id", "updated_at", "created_at"}).
		AddRow(mockProducts[0].Id, mockProducts[0].Style, mockProducts[0].GenderId, mockProducts[0].SizeId, mockProducts[0].ColorId,
			priceZero, mockProducts[0].PatternId, mockProducts[0].UpdatedAt, mockProducts[0].CreatedAt).
		AddRow(mockProducts[1].Id, mockProducts[1].Style, mockProducts[1].GenderId, mockProducts[1].SizeId, mockProducts[1].ColorId,
			priceOne, mockProducts[1].PatternId, mockProducts[1].UpdatedAt, mockProducts[1].CreatedAt)

	query := "select * from products as p where 1=1 "

	mock.ExpectQuery(query).WillReturnRows(rows)
	p := productMysqlRepo.NewProductRepository(db)

	list, _, err := p.GetAllProducts(context.TODO(), "", "", "", "", "")
	assert.NoError(t, err)
	assert.Len(t, list, 2)

}
