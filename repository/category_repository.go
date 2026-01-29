package repository

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/kenzchiro/desksetup-link-api/domain"
)

type CategoryRepository struct {
	db      *sqlx.DB
	timeout time.Duration
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		db:      db,
		timeout: 5 * time.Second,
	}
}

func (r *CategoryRepository) List(ctx context.Context) ([]domain.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	var categories []domain.Category
	err := r.db.SelectContext(ctx, &categories, `SELECT id, name, description, seq FROM categories ORDER BY seq, id`)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
