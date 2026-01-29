package product

import (
	"time"

	"github.com/jmoiron/sqlx"
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
