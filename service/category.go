package service

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]domain.Category, error)
}

type CategoryService struct {
	repo CategoryRepository
}

func NewCategoryService(repo CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) List(ctx context.Context) ([]domain.Category, error) {
	return s.repo.List(ctx)
}
