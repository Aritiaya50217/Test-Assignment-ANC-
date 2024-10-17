package handler

import (
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/Test-Assignment-ANC/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(ctx *gin.Context) {
	genderId, _ := strconv.Atoi(ctx.Query("gender_id"))
	style := ctx.Query("style")
	sizeId, _ := strconv.Atoi(ctx.Query("size_id"))
	offset, _ := strconv.Atoi(ctx.Query("offset"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	perPage := limit * (offset - 1)

	products, total, _ := models.GetAllProducts(genderId, style, sizeId, offset, limit, perPage)
	// TODO : check product id

	result := map[string]interface{}{
		"products": products,
		"total":    total,
		"per_page": perPage,
	}
	ctx.JSON(http.StatusOK, result)
}

func GetProduct(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	// TODO : query
	// product, _ := models.GetProductById(id)
	// if err := h.db.Find(&product, id).Error; err != nil {
	// 	ctx.Status(http.StatusNotFound)
	// 	return
	// }

	// mock data
	result := map[string]interface{}{
		"id":     id,
		"gender": "Men",
		"style":  "Plain color / Red",
		"size":   "XS",
		"price":  "100",
	}
	ctx.JSONP(http.StatusOK, result)

}
