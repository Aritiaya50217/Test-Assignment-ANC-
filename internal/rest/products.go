package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
	"github.com/Aritiaya50217/Test-Assignment-ANC/utils"
	"github.com/labstack/echo/v4"
)

// ProductService represent the product's usecase
type ProductService interface {
	GetProductById(ctx context.Context, id int64) (domain.Product, error)
	GetAllProducts(ctx context.Context, genderId int, style string, sizeId, offset, limit int) ([]domain.Product, int, error)
}

// ProductHandler represent the http handler for product
type ProductHandler struct {
	Service ProductService
}

// NewProductHandler will initialize the product resourse endpoint
func NewProductHandler(e *echo.Echo, svc ProductService) {
	handler := &ProductHandler{
		Service: svc,
	}
	e.GET("/products/:id", handler.GetProductById)
	e.GET("/products", handler.GetAllProducts)
}

func (p *ProductHandler) GetProductById(c echo.Context) error {
	idP, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	id := int64(idP)
	ctx := c.Request().Context()

	product, err := p.Service.GetProductById(ctx, id)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, product)
}

func (p *ProductHandler) GetAllProducts(c echo.Context) error {
	
	genderId, _ := strconv.Atoi(c.QueryParam("gender_id"))
	style := c.QueryParam("style")
	sizeId, _ := strconv.Atoi(c.QueryParam("size_id"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}

	ctx := c.Request().Context()
	products, total, err := p.Service.GetAllProducts(ctx, genderId, style, sizeId, offset, limit)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}
	result := map[string]interface{}{
		"products": &products,
		"total":    total,
	}
	return c.JSON(http.StatusOK, result)
}
