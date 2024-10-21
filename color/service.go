package color

import (
	"context"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

// ColorRepository represent the color's repository contract
type ColorRepository interface {
	GetColorById(ctx context.Context, id int64) (domain.Color, error)
}

type Service struct {
	colorRepo ColorRepository
}

// NewService will create a new color service object
func NewService(c ColorRepository) *Service {
	return &Service{
		colorRepo: c,
	}
}

func (s *Service) GetColorById(ctx context.Context, id int64) (res domain.Color, err error) {
	res, err = s.colorRepo.GetColorById(ctx, id)
	if err != nil {
		return
	}
	resColor, err := s.colorRepo.GetColorById(ctx, res.Id)
	if err != nil {
		return domain.Color{}, err
	}
	return resColor, err
}
