package highlight

import (
	"context"
	"time"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *HighlightRepository) Update(ctx context.Context, id int64, h domain.Highlight) (*domain.Highlight, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	h.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, `
		UPDATE highlights
		SET product_id = $1, priority = $2, end_date = $3, updated_at = $4
		WHERE id = $5
	`,
		h.ProductID,
		h.Priority,
		h.EndDate,
		h.UpdatedAt,
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

	h.ID = id
	return &h, true, nil
}
