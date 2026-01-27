package service

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

// HighlightRepository defines the persistence port for highlights.
type HighlightRepository interface {
	List(ctx context.Context) ([]domain.Highlight, error)
	Get(ctx context.Context, id int64) (*domain.Highlight, bool, error)
	GetByProductID(ctx context.Context, productID int64) (*domain.Highlight, bool, error)
	Create(ctx context.Context, h domain.Highlight) (domain.Highlight, error)
	Update(ctx context.Context, id int64, h domain.Highlight) (*domain.Highlight, bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

// HighlightService coordinates business logic for highlights.
type HighlightService struct {
	repo HighlightRepository
}

func NewHighlightService(repo HighlightRepository) *HighlightService {
	return &HighlightService{repo: repo}
}

func (s *HighlightService) List(ctx context.Context) ([]domain.Highlight, error) {
	return s.repo.List(ctx)
}

func (s *HighlightService) Get(ctx context.Context, id int64) (*domain.Highlight, bool, error) {
	return s.repo.Get(ctx, id)
}

func (s *HighlightService) GetByProductID(ctx context.Context, productID int64) (*domain.Highlight, bool, error) {
	return s.repo.GetByProductID(ctx, productID)
}

func (s *HighlightService) Create(ctx context.Context, h domain.Highlight) (domain.Highlight, error) {
	return s.repo.Create(ctx, h)
}

func (s *HighlightService) Update(ctx context.Context, id int64, h domain.Highlight) (*domain.Highlight, bool, error) {
	return s.repo.Update(ctx, id, h)
}

func (s *HighlightService) Delete(ctx context.Context, id int64) (bool, error) {
	return s.repo.Delete(ctx, id)
}
