package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kenzchiro/desksetup-link-api/domain"
)

type SubItemRepository struct {
	db      *sqlx.DB
	timeout time.Duration
}

func NewSubItemRepository(db *sqlx.DB) *SubItemRepository {
	return &SubItemRepository{
		db:      db,
		timeout: 5 * time.Second,
	}
}

func (r *SubItemRepository) ListByProductID(ctx context.Context, productID int64) ([]domain.SubItem, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var dbSubItems []struct {
		domain.SubItem
		CategoryJSON string `db:"category"`
	}

	err := r.db.SelectContext(ctx, &dbSubItems, `
		SELECT id, product_id, title, subtitle, brand, img, category, description, code,
		       shopee_link, tiktok_link, lazada_link, other_link, display_order,
		       created_at, updated_at
		FROM sub_items
		WHERE product_id = $1
		ORDER BY display_order ASC, id ASC
	`, productID)

	if err != nil {
		return nil, err
	}

	subItems := make([]domain.SubItem, 0, len(dbSubItems))
	for _, dbsi := range dbSubItems {
		si := dbsi.SubItem
		if dbsi.CategoryJSON != "" {
			_ = json.Unmarshal([]byte(dbsi.CategoryJSON), &si.Category)
		}
		si.ToLinks()
		subItems = append(subItems, si)
	}

	return subItems, nil
}

func (r *SubItemRepository) Get(ctx context.Context, id int64) (*domain.SubItem, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var dbSubItem struct {
		domain.SubItem
		CategoryJSON string `db:"category"`
	}

	err := r.db.GetContext(ctx, &dbSubItem, `
		SELECT id, product_id, title, subtitle, brand, img, category, description, code,
		       shopee_link, tiktok_link, lazada_link, other_link, display_order,
		       created_at, updated_at
		FROM sub_items
		WHERE id = $1
		LIMIT 1
	`, id)

	if err == sql.ErrNoRows {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	si := dbSubItem.SubItem
	if dbSubItem.CategoryJSON != "" {
		_ = json.Unmarshal([]byte(dbSubItem.CategoryJSON), &si.Category)
	}
	si.ToLinks()

	return &si, true, nil
}

func (r *SubItemRepository) Create(ctx context.Context, si domain.SubItem) (domain.SubItem, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	categoryJSON, err := json.Marshal(si.Category)
	if err != nil {
		return domain.SubItem{}, err
	}

	now := time.Now()
	si.CreatedAt = now
	si.UpdatedAt = now

	err = r.db.QueryRowContext(ctx, `
		INSERT INTO sub_items (product_id, title, subtitle, brand, img, category, description, code,
		                       shopee_link, tiktok_link, lazada_link, other_link, display_order,
		                       created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		RETURNING id
	`,
		si.ProductID,
		si.Title,
		si.Subtitle,
		si.Brand,
		si.Img,
		string(categoryJSON),
		si.Description,
		si.Code,
		si.ShopeeLink,
		si.TiktokLink,
		si.LazadaLink,
		si.OtherLink,
		si.DisplayOrder,
		si.CreatedAt,
		si.UpdatedAt,
	).Scan(&si.ID)

	if err != nil {
		return domain.SubItem{}, err
	}

	si.ToLinks()
	return si, nil
}

func (r *SubItemRepository) Update(ctx context.Context, id int64, si domain.SubItem) (*domain.SubItem, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	categoryJSON, err := json.Marshal(si.Category)
	if err != nil {
		return nil, false, err
	}

	si.UpdatedAt = time.Now()

	result, err := r.db.ExecContext(ctx, `
		UPDATE sub_items
		SET title = $1, subtitle = $2, brand = $3, img = $4, category = $5, description = $6,
		    code = $7, shopee_link = $8, tiktok_link = $9, lazada_link = $10, other_link = $11,
		    display_order = $12, updated_at = $13
		WHERE id = $14
	`,
		si.Title,
		si.Subtitle,
		si.Brand,
		si.Img,
		string(categoryJSON),
		si.Description,
		si.Code,
		si.ShopeeLink,
		si.TiktokLink,
		si.LazadaLink,
		si.OtherLink,
		si.DisplayOrder,
		si.UpdatedAt,
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

	si.ID = id
	si.ToLinks()
	return &si, true, nil
}

func (r *SubItemRepository) Delete(ctx context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	result, err := r.db.ExecContext(ctx, `
		DELETE FROM sub_items WHERE id = $1
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
