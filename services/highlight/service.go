package highlight

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
	"github.com/kenzchiro/desksetup-link-api/services/product"
)

type HighlightRepository interface {
	List(ctx context.Context) ([]domain.Highlight, error)
	Create(ctx context.Context, h domain.Highlight) (domain.Highlight, error)
	Update(ctx context.Context, id int64, h domain.Highlight) (*domain.Highlight, bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type HighlightService struct {
	highlightRepo HighlightRepository
	productRepo   product.ProductRepository
}

func NewHighlightService(highlightRepo HighlightRepository, productRepo product.ProductRepository) *HighlightService {
	return &HighlightService{highlightRepo: highlightRepo, productRepo: productRepo}
}
