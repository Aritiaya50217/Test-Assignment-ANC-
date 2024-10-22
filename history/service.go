package history

import (
	"context"

	"github.com/Aritiaya50217/Test-Assignment-ANC/domain"
)

// hisrotyRepository represent the history's repository contract
type HistoryRepository interface {
	InsertHistory(ctx context.Context, o *domain.History) error
}

type Service struct {
	historyRepo HistoryRepository
}

// NewService will create a new history service object
func NewService(h HistoryRepository) *Service {
	return &Service{
		historyRepo: h,
	}
}

func (s *Service) InsertHistory(ctx context.Context, h *domain.History) (err error) {
	err = s.historyRepo.InsertHistory(ctx, h)
	if err != nil {
		return
	}
	return
}
