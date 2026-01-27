package domain

import (
	"time"
)

// Product represents a product and its purchase links.
type Product struct {
	ID          int64             `db:"id" json:"id"`
	Title       string            `db:"title" json:"title"`
	Brand       *string           `db:"brand" json:"brand"`
	Img         *string           `db:"img" json:"img"`
	Category    []string          `db:"category" json:"category"`
	Description *string           `db:"description" json:"description"`
	Code        *string           `db:"code" json:"code"`
	Tag         *string           `db:"tag" json:"tag"`
	Links       map[string]string `db:"links" json:"links"`
	CreatedAt   time.Time         `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time         `db:"updated_at" json:"updated_at"`
}

// SubItem represents a product variant or grouped item.
type SubItem struct {
	ID           int64     `db:"id" json:"id"`
	ProductID    int64     `db:"product_id" json:"product_id"`
	Title        string    `db:"title" json:"title"`
	Subtitle     *string   `db:"subtitle" json:"subtitle"`
	Brand        *string   `db:"brand" json:"brand"`
	Img          *string   `db:"img" json:"img"`
	Category     []string  `db:"category" json:"category"`
	Description  *string   `db:"description" json:"description"`
	Code         *string   `db:"code" json:"code"`
	ShopeeLink   *string   `db:"shopee_link" json:"shopee_link"`
	TiktokLink   *string   `db:"tiktok_link" json:"tiktok_link"`
	LazadaLink   *string   `db:"lazada_link" json:"lazada_link"`
	OtherLink    *string   `db:"other_link" json:"other_link"`
	DisplayOrder int16     `db:"display_order" json:"display_order"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

// Highlight represents a featured product.
type Highlight struct {
	ID        int64      `db:"id" json:"id"`
	ProductID int64      `db:"product_id" json:"product_id"`
	Priority  int16      `db:"priority" json:"priority"`
	EndDate   *time.Time `db:"end_date" json:"end_date"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
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
