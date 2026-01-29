package category

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (s *CategoryService) List(ctx context.Context) ([]domain.Category, error) {
	return s.categoryRepo.List(ctx)
}
