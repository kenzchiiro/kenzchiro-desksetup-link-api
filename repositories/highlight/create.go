package highlight

import (
	"context"
	"time"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *HighlightRepository) Create(ctx context.Context, h domain.Highlight) (domain.Highlight, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	now := time.Now()
	h.CreatedAt = now
	h.UpdatedAt = now

	err := r.db.QueryRowContext(ctx, `
		INSERT INTO highlights (product_id, priority, end_date, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`,
		h.ProductID,
		h.Priority,
		h.EndDate,
		h.CreatedAt,
		h.UpdatedAt,
	).Scan(&h.ID)

	if err != nil {
		return domain.Highlight{}, err
	}

	return h, nil
}
