package rest

import (
	"context"
	"net/http"
	"time"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
	"github.com/Aritiaya50217/Test-Assignment-ANC/utils"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// OrderService represent the order's usecase
type OrderService interface {
	GetAllOrders(ctx context.Context, startDate, endDate, statusId, offset, limit string) ([]domain.Order, int, error)
	InsertOrder(context.Context, *domain.Order) error
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
	e.POST("/orders", handler.InsertOrder)
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

func isRequestValid(o *domain.Order) (bool, error) {
	validate := validator.New()
	err := validate.Struct(o)
	if err != nil {
		return false, err
	}
	return true, nil
}

// InsertOrder will store the order by given request body
func (o *OrderHandler) InsertOrder(c echo.Context) (err error) {
	var orders domain.Orders
	err = c.Bind(&orders)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.ResponseError{
			Message: err.Error(),
		})
	}
	var data domain.Order
	var ok bool
	if ok, err = isRequestValid(&data); !ok {
		return c.JSON(http.StatusBadRequest, utils.ResponseError{
			Message: err.Error(),
		})
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	ctx := c.Request().Context()

	for _, val := range orders.Orders.Products {
		data.ProductId = val.ProductId
		data.Amount = val.Amount
		data.UserId = orders.Orders.UserId
		data.Address = orders.Orders.Address
		data.CreatedAt = now
		data.UpdatedAt = now

		// insert to table orders
		err = o.Service.InsertOrder(ctx, &data)
		if err != nil {
			return c.JSON(utils.GetStatusCode(err), utils.ResponseError{
				Message: err.Error(),
			})
		}
	}

	return c.JSON(http.StatusCreated, nil)
}
