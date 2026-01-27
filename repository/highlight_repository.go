package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kenzchiro/desksetup-link-api/domain"
)

type HighlightRepository struct {
	db      *sqlx.DB
	timeout time.Duration
}

func NewHighlightRepository(db *sqlx.DB) *HighlightRepository {
	return &HighlightRepository{
		db:      db,
		timeout: 5 * time.Second,
	}
}

func (r *HighlightRepository) List(ctx context.Context) ([]domain.Highlight, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var highlights []domain.Highlight
	err := r.db.SelectContext(ctx, &highlights, `
		SELECT id, product_id, priority, end_date, created_at, updated_at
		FROM highlights
		ORDER BY priority DESC, created_at DESC
	`)
	if err != nil {
		return nil, err
	}

	return highlights, nil
}

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

func (r *HighlightRepository) GetByProductID(ctx context.Context, productID int64) (*domain.Highlight, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var h domain.Highlight
	err := r.db.GetContext(ctx, &h, `
		SELECT id, product_id, priority, end_date, created_at, updated_at
		FROM highlights
		WHERE product_id = $1
		LIMIT 1
	`, productID)

	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	return &h, true, nil
}

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

func (r *HighlightRepository) Delete(ctx context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	result, err := r.db.ExecContext(ctx, `
		DELETE FROM highlights WHERE id = $1
	`, id)

	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected > 0, nil
}
