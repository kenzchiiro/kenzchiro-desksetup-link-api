package product

import (
	"context"
	"encoding/json"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *ProductRepository) List(ctx context.Context) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var dbProducts []struct {
		domain.Product
		CategoryJSON string `db:"category"`
		LinksJSON    string `db:"links"`
	}

	err := r.db.SelectContext(ctx, &dbProducts, `
		SELECT id, title, category, brand, img, tag, description, code, links, created_at, updated_at, parent_product
		FROM products
		WHERE deleted_at IS NULL
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}

	products := make([]domain.Product, 0, len(dbProducts))
	for _, dbp := range dbProducts {
		p := dbp.Product
		if dbp.CategoryJSON != "" {
			_ = json.Unmarshal([]byte(dbp.CategoryJSON), &p.Category)
		}
		if dbp.LinksJSON != "" {
			_ = json.Unmarshal([]byte(dbp.LinksJSON), &p.Links)
		}
		products = append(products, p)
	}

	return products, nil
}
