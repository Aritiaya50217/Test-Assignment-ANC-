package product

import (
	"context"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

// ProductRepository represent the product's repository contract
type ProductRepository interface {
	GetProductById(ctx context.Context, id int64) (domain.Product, error)
	GetAllProducts(ctx context.Context, genderId int, style string, sizeId, offset, limit int) ([]domain.Product, int, error)
}

type Service struct {
	productRepo ProductRepository
}

// NewProduct will create a new product service object
func NewService(p ProductRepository) *Service {
	return &Service{
		productRepo: p,
	}
}

func (s *Service) GetProductById(ctx context.Context, id int64) (res domain.Product, err error) {
	res, err = s.productRepo.GetProductById(ctx, id)
	if err != nil {
		return
	}

	resProduct, err := s.productRepo.GetProductById(ctx, res.Id)
	if err != nil {
		return domain.Product{}, err
	}
	return resProduct, err
}

func (s *Service) GetAllProducts(ctx context.Context, genderId int, style string, sizeId, offset, limit int) (res []domain.Product, total int, err error) {
	res, _, err = s.productRepo.GetAllProducts(ctx, genderId, style, sizeId, offset, limit)
	if err != nil {
		return
	}
	resProducts, total, err := s.productRepo.GetAllProducts(ctx, genderId, style, sizeId, offset, limit)
	if err != nil {
		return []domain.Product{}, 0, err
	}
	return resProducts, total, err
}
