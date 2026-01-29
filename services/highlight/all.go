package highlight

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
	"github.com/kenzchiro/desksetup-link-api/services/product"
)

func (s *HighlightService) List(ctx context.Context) ([]domain.Highlight, error) {
	highlights, err := s.highlightRepo.List(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch all products and build sub product mapping
	prodSvc := product.NewProductService(s.productRepo)
	allProducts, err := prodSvc.List(ctx)
	if err != nil {
		return highlights, nil // fallback: return highlights without sub_products
	}
	productMap := make(map[int64]domain.Product)
	for _, p := range allProducts {
		productMap[p.ID] = p
	}

	for i := range highlights {
		if highlights[i].Product != nil {
			if prod, ok := productMap[highlights[i].Product.ID]; ok {
				highlights[i].Product.SubProducts = prod.SubProducts
			}
		}
	}
	return highlights, nil
}
