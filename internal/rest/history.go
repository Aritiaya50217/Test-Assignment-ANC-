package rest

import (
	"github.com/labstack/echo/v4"
)

// HistoryService represent the history's usecase
type HistoryService interface {
}

// HistoryHandler represent the http handler for history
type HistoryHandler struct {
	Service HistoryService
}

// NewHistoryHandler will initialize the history resourse endpoint
func NewHistoryHandler(e *echo.Echo, svc HistoryService) {
	handler := &HistoryHandler{
		Service: svc,
	}
	_ = handler
}
