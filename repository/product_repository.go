package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kenzchiro/desksetup-link-api/domain"
)

type ProductRepository struct {
	db      *sqlx.DB
	timeout time.Duration
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db:      db,
		timeout: 5 * time.Second,
	}
}

func (r *ProductRepository) List(ctx context.Context) ([]domain.Product, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var dbProducts []struct {
		domain.Product
		CategoryJSON string `db:"category"`
		LinksJSON    string `db:"links"`
	}

	err := r.db.SelectContext(ctx, &dbProducts, `
		SELECT id, title, category, brand, img, tag, description, code, links, created_at, updated_at
		FROM products
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}

	products := make([]domain.Product, 0, len(dbProducts))
	for _, dbp := range dbProducts {
		p := dbp.Product
		if dbp.CategoryJSON != "" {
			_ = json.Unmarshal([]byte(dbp.CategoryJSON), &p.Category)
		}
		if dbp.LinksJSON != "" {
			_ = json.Unmarshal([]byte(dbp.LinksJSON), &p.Links)
		}

		// Fetch sub_items for this product
		subItems, _ := r.getSubItems(ctx, p.ID)
		p.SubItems = subItems

		products = append(products, p)
	}

	return products, nil
}

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

	// Fetch sub_items for this product
	subItems, _ := r.getSubItems(ctx, p.ID)
	p.SubItems = subItems

	return &p, true, nil
}

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

// getSubItems fetches sub_items for a product and converts individual link fields to Links map
func (r *ProductRepository) getSubItems(ctx context.Context, productID int64) ([]domain.SubItem, error) {
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
		// Convert individual link fields to Links map
		si.ToLinks()
		subItems = append(subItems, si)
	}

	return subItems, nil
}

func (r *ProductRepository) Delete(ctx context.Context, id int64) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	result, err := r.db.ExecContext(ctx, `
		DELETE FROM products WHERE id = $1
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
