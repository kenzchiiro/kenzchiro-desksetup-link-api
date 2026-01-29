package product

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (s *ProductService) Get(ctx context.Context, id int64) (*domain.Product, bool, error) {
	return s.productRepo.Get(ctx, id)
}
