package product

import (
	"context"
	"encoding/json"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *ProductRepository) Update(ctx context.Context, id int64, p domain.Product) (*domain.Product, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	categoryJSON, err := json.Marshal(p.Category)
	if err != nil {
		return nil, false, err
	}

	linksJSON, err := json.Marshal(p.Links)
	if err != nil {
		return nil, false, err
	}

	result, err := r.db.ExecContext(ctx, `
		UPDATE products
		SET title = $1, category = $2, brand = $3, img = $4, tag = $5, description = $6, code = $7, links = $8
		WHERE id = $9
	`,
		p.Title,
		string(categoryJSON),
		p.Brand,
		p.Img,
		p.Tag,
		p.Description,
		p.Code,
		string(linksJSON),
		id,
	)
	if err != nil {
		return nil, false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return nil, false, err
	}
	if affected == 0 {
		return nil, false, nil
	}

	p.ID = id
	return &p, true, nil
}
