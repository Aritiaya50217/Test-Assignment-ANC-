package rest

import (
	"context"
	"net/http"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
	"github.com/Aritiaya50217/Test-Assignment-ANC/utils"
	"github.com/labstack/echo/v4"
)

// OrderService represent the order's usecase
type OrderService interface {
	GetAllOrders(ctx context.Context, startDate, endDate, statusId, offset, limit string) ([]domain.Order, int, error)
}

// OrderHandler represent the http handler for order
type OrderHandler struct {
	Service OrderService
}

// NewOrderHandler will initialize the order resourse endpoint
func NewOrderHandler(e *echo.Echo, svc OrderService) {
	handler := &OrderHandler{
		Service: svc,
	}
	e.GET("/orders", handler.GetAllOrders)
}

func (o *OrderHandler) GetAllOrders(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	statusId := c.QueryParam("status_id")
	offset := c.QueryParam("offset")
	limit := c.QueryParam("limit")
	if limit == "" {
		limit = "10"
	}

	ctx := c.Request().Context()
	orders, total, err := o.Service.GetAllOrders(ctx, startDate, endDate, statusId, offset, limit)
	if err != nil {
		return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}
	result := map[string]interface{}{
		"orders": orders,
		"total":  total,
	}

	return c.JSON(http.StatusOK, result)
}
