package order

import (
	"context"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

// orderRepository represent the order's repository contract
type OrderRepository interface {
	GetAllOrders(ctx context.Context, startDate, endDate, statusId, offset, limit string) ([]domain.Order, int, error)
}

type Service struct {
	orderRepo OrderRepository
}

// NewOrder will create a new order service object
func NewService(o OrderRepository) *Service {
	return &Service{
		orderRepo: o,
	}
}

func (s *Service) GetAllOrders(ctx context.Context, startDate, endDate, statusId, offset, limit string) (res []domain.Order, total int, err error) {
	res, _, err = s.orderRepo.GetAllOrders(ctx, startDate, endDate, statusId, offset, limit)
	if err != nil {
		return
	}
	resOrders, total, err := s.orderRepo.GetAllOrders(ctx, startDate, endDate, statusId, offset, limit )
	if err != nil {
		return []domain.Order{}, 0, err
	}
	return resOrders, total, err
}
