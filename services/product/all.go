package product

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (s *ProductService) List(ctx context.Context) ([]domain.Product, error) {
	products, err := s.productRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return mapGroupItems(products), nil
}

func mapGroupItems(products []domain.Product) []domain.Product {
	productsMap := make(map[int64]*domain.Product)
	for i := range products {
		productsMap[products[i].ID] = &products[i]
	}
	for id, product := range productsMap {
		if product.ParentProduct != nil {
			parent, ok := productsMap[*product.ParentProduct]
			if ok {
				productsMap[parent.ID].SubProducts = append(productsMap[parent.ID].SubProducts, *productsMap[id])
			}
		}
	}

	parents := []domain.Product{}
	for _, product := range productsMap {
		if product.ParentProduct == nil {
			parents = append(parents, *product)
		}
	}
	return parents
}
