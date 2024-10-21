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
