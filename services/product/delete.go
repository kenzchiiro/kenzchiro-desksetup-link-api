package product

import (
	"context"
)

func (s *ProductService) Delete(ctx context.Context, id int64) (bool, error) {
	return s.productRepo.Delete(ctx, id)
}
