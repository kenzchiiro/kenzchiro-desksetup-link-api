package product

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (s *ProductService) Create(ctx context.Context, p domain.Product) (domain.Product, error) {
	if err := p.Validate(); err != nil {
		return domain.Product{}, err
	}
	return s.productRepo.Create(ctx, p)
}
