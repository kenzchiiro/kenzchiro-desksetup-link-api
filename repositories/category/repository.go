package category

import (
	"time"

	"github.com/jmoiron/sqlx"
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
