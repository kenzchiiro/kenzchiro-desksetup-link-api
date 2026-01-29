package domain

import (
	"time"
)

// Product represents a product and its purchase links.
type Product struct {
	ID            int64             `db:"id" json:"-"`
	Title         string            `db:"title" json:"title"`
	Brand         *string           `db:"brand" json:"brand"`
	Img           *string           `db:"img" json:"img"`
	Category      []string          `db:"category" json:"category"`
	Description   *string           `db:"description" json:"description"`
	Code          *string           `db:"code" json:"code"`
	Tag           *string           `db:"tag" json:"tag"`
	Links         map[string]string `db:"links" json:"links"`
	CreatedAt     time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time         `db:"updated_at" json:"updated_at"`
	ParentProduct *int64            `db:"parent_product" json:"parent_product,omitempty"`
	SubProducts   []Product         `db:"-" json:"sub_products,omitempty"`
}

// Highlight represents a featured product.
type Highlight struct {
	ID        int64      `db:"id" json:"-"`
	ProductID int64      `db:"product_id" json:"-"`
	Priority  int16      `db:"priority" json:"priority"`
	EndDate   *time.Time `db:"end_date" json:"end_date"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	Product   *Product   `db:"-" json:"product,omitempty"`
}

// Validate checks minimal required fields.
func (p *Product) Validate() error {
	if p.Title == "" {
		return ErrInvalidProductTitle
	}
	if p.Links == nil {
		p.Links = map[string]string{}
	}
	if p.Category == nil {
		p.Category = []string{}
	}
	return nil
}
