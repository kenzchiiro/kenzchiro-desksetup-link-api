package product

import (
	"context"
	"encoding/json"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *ProductRepository) Create(ctx context.Context, p domain.Product) (domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	categoryJSON, err := json.Marshal(p.Category)
	if err != nil {
		return domain.Product{}, err
	}

	linksJSON, err := json.Marshal(p.Links)
	if err != nil {
		return domain.Product{}, err
	}

	err = r.db.QueryRowContext(ctx, `
		INSERT INTO products (title, category, brand, img, tag, description, code, links)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`,
		p.Title,
		string(categoryJSON),
		p.Brand,
		p.Img,
		p.Tag,
		p.Description,
		p.Code,
		string(linksJSON),
	).Scan(&p.ID)
	if err != nil {
		return domain.Product{}, err
	}

	return p, nil
}
