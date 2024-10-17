package routes

import (
	"net/http"

	"github.com/Aritiaya50217/Test-Assignment-ANC/handler"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	productGroup := r.Group("/product")
	productGroup.GET("/list", handler.GetProducts)
	productGroup.GET("/:id", handler.GetProduct)

	orderGroup := r.Group("/order")
	orderGroup.POST("/", handler.CreateOrder)
	orderGroup.GET("/list", handler.GetOrders)

	return r
}
