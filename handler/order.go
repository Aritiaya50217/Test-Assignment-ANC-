package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/Test-Assignment-ANC/errs"
	"github.com/Aritiaya50217/Test-Assignment-ANC/models"
	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	Orders []Products `json:"orders"`
}

type Products struct {
	Product int `json:"product_id"`
	User    int `json:"user_id"`
}

func GetOrders(ctx *gin.Context) {
	startDate := ctx.Query("start_date")
	endDate := ctx.Query("end_date")
	statusId, _ := strconv.Atoi(ctx.Query("status_id"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	perPage := limit * (offset - 1)

	orders, _, _ := models.GetOrders(startDate, endDate, statusId, offset, limit, perPage)
	items := make([]map[string]interface{}, 0)
	for _, order := range orders {
		items = append(items, genItemOrder(order))
	}
	// mock
	items = append(items, map[string]interface{}{
		"id":         1,
		"product_id": 1,
		"user_id":    1,
	}, map[string]interface{}{
		"id":         2,
		"product_id": 2,
		"user_id":    1,
	})
	result := map[string]interface{}{
		"orders": items,
		"total":  len(items),
	}

	ctx.JSON(http.StatusOK, result)
}

func CreateOrder(ctx *gin.Context) {
	var orders OrderRequest
	err := ctx.BindJSON(&orders)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errs.NewUnexpectedError().Error())
		return
	}
	for _, order := range orders.Orders {
		fmt.Println(order)
		// TODO : create new order

	}
	ctx.JSON(http.StatusOK, "success")
}

func genItemOrder(order models.Order) (result map[string]interface{}) {
	if order.Id != 0 {
		result = map[string]interface{}{
			"id":         order.Id,
			"product_id": order.Product.Id,
			"user_id":    order.User.Id,
		}
	}
	return
}
