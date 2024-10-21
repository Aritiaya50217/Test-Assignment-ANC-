package rest

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
	"github.com/Aritiaya50217/Test-Assignment-ANC/utils"
	"github.com/labstack/echo/v4"
)

// ColorService represent the color's usecase
type ColorService interface {
	GetColorById(ctx context.Context, id int64) (domain.Color, error)
}

// ColorHandler represent the http handler for color
type ColorHandler struct {
	Service ColorService
}

// NewColorHandler will initialize the color resoures endpoint
func NewColorHandler(e *echo.Echo, svc ColorService) {
	handler := &ColorHandler{
		Service: svc,
	}
	e.GET("/colors/:id", handler.GetColorById)
}

func (c *ColorHandler) GetColorById(e echo.Context) error {
	idP, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return e.JSON(http.StatusNotFound, domain.ErrNotFound.Error())
	}
	id := int64(idP)
	ctx := e.Request().Context()

	color, err := c.Service.GetColorById(ctx, id)
	if err != nil {
		return e.JSON(utils.GetStatusCode(err), utils.ResponseError{
			Message: err.Error(),
		})
	}
	return e.JSON(http.StatusOK, color)
}
