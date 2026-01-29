package repository

import (
	"context"
	"database/sql"
	"encoding/json"
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

	type highlightWithProduct struct {
		domain.Highlight
		// Product fields for join
		ProductID     int64      `db:"p_id"`
		Title         string     `db:"p_title"`
		Brand         *string    `db:"p_brand"`
		Img           *string    `db:"p_img"`
		CategoryJSON  string     `db:"p_category"`
		Description   *string    `db:"p_description"`
		Code          *string    `db:"p_code"`
		Tag           *string    `db:"p_tag"`
		LinksJSON     string     `db:"p_links"`
		CreatedAtProd *time.Time `db:"p_created_at"`
		UpdatedAtProd *time.Time `db:"p_updated_at"`
	}

	var rows []highlightWithProduct
	err := r.db.SelectContext(ctx, &rows, `
		       SELECT h.id, h.product_id, h.priority, h.end_date, h.created_at, h.updated_at,
			      p.id as p_id, p.title as p_title, p.brand as p_brand, p.img as p_img, p.category as p_category, p.description as p_description, p.code as p_code, p.tag as p_tag, p.links as p_links, p.created_at as p_created_at, p.updated_at as p_updated_at
		       FROM highlights h
		       JOIN products p ON h.product_id = p.id
		       ORDER BY h.priority DESC, h.created_at DESC
	       `)
	if err != nil {
		return nil, err
	}

	highlights := make([]domain.Highlight, 0, len(rows))
	for _, row := range rows {
		prod := &domain.Product{
			ID:          row.ProductID,
			Title:       row.Title,
			Brand:       row.Brand,
			Img:         row.Img,
			Description: row.Description,
			Code:        row.Code,
			Tag:         row.Tag,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		}
		if row.CreatedAtProd != nil {
			prod.CreatedAt = *row.CreatedAtProd
		}
		if row.UpdatedAtProd != nil {
			prod.UpdatedAt = *row.UpdatedAtProd
		}
		if row.CategoryJSON != "" {
			_ = json.Unmarshal([]byte(row.CategoryJSON), &prod.Category)
		}
		if row.LinksJSON != "" {
			_ = json.Unmarshal([]byte(row.LinksJSON), &prod.Links)
		}
		h := row.Highlight
		h.Product = prod
		highlights = append(highlights, h)
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
