package highlight

import (
	"context"
	"database/sql"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *HighlightRepository) Get(ctx context.Context, id int64) (*domain.Highlight, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var h domain.Highlight
	err := r.db.GetContext(ctx, &h, `
		SELECT id, product_id, priority, end_date, created_at, updated_at
		FROM highlights
		WHERE id = $1
		LIMIT 1
	`, id)

	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	return &h, true, nil
}
