package category

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *CategoryRepository) List(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var categories []domain.Category
	err := r.db.SelectContext(ctx, &categories, `SELECT id, name, description, seq, icon FROM categories ORDER BY seq, id`)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
