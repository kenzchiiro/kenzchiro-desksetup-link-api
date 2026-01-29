package product

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

type ProductRepository interface {
	List(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int64) (*domain.Product, bool, error)
	Create(ctx context.Context, p domain.Product) (domain.Product, error)
	Update(ctx context.Context, id int64, p domain.Product) (*domain.Product, bool, error)
	Delete(ctx context.Context, id int64) (bool, error)
}

type ProductService struct {
	productRepo ProductRepository
}

func NewProductService(productRepo ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}
