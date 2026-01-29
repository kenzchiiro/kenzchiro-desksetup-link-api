package category

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]domain.Category, error)
}

type CategoryService struct {
	categoryRepo CategoryRepository
}

func NewCategoryService(categoryRepo CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepo: categoryRepo}
}
