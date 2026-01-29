package highlight

import (
	"time"

	"github.com/jmoiron/sqlx"
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
