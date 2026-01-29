package product

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *ProductRepository) Get(ctx context.Context, id int64) (*domain.Product, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var dbProduct struct {
		domain.Product
		CategoryJSON string `db:"category"`
		LinksJSON    string `db:"links"`
	}

	err := r.db.GetContext(ctx, &dbProduct, `
		SELECT id, title, category, brand, img, tag, description, code, links, created_at, updated_at
		FROM products
		WHERE id = $1
		LIMIT 1
	`, id)

	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	p := dbProduct.Product
	if dbProduct.CategoryJSON != "" {
		_ = json.Unmarshal([]byte(dbProduct.CategoryJSON), &p.Category)
	}
	if dbProduct.LinksJSON != "" {
		_ = json.Unmarshal([]byte(dbProduct.LinksJSON), &p.Links)
	}

	return &p, true, nil
}
