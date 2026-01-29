package product

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (s *ProductService) Update(ctx context.Context, id int64, p domain.Product) (*domain.Product, bool, error) {
	if err := p.Validate(); err != nil {
		return nil, false, err
	}
	return s.productRepo.Update(ctx, id, p)
}
