package service

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

// ProductRepository defines the persistence port.
type ProductRepository interface {
	List(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int64) (*domain.Product, bool, error)
	Create(ctx context.Context, p domain.Product) (domain.Product, error)
	Update(ctx context.Context, id int64, p domain.Product) (*domain.Product, bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

// ProductService coordinates business logic.
type ProductService struct {
	repo ProductRepository
}

func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) List(ctx context.Context) ([]domain.Product, error) {
	return s.repo.List(ctx)
}

func (s *ProductService) Get(ctx context.Context, id int64) (*domain.Product, bool, error) {
	return s.repo.Get(ctx, id)
}

func (s *ProductService) Create(ctx context.Context, p domain.Product) (domain.Product, error) {
	if err := p.Validate(); err != nil {
		return domain.Product{}, err
	}
	return s.repo.Create(ctx, p)
}

func (s *ProductService) Update(ctx context.Context, id int64, p domain.Product) (*domain.Product, bool, error) {
	if err := p.Validate(); err != nil {
		return nil, false, err
	}
	return s.repo.Update(ctx, id, p)
}

func (s *ProductService) Delete(ctx context.Context, id int64) (bool, error) {
	return s.repo.Delete(ctx, id)
}
