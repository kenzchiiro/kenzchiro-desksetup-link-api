package domain

import (
	"time"
)

// Product represents a product and its purchase links.
type Product struct {
	ID          int64             `db:"id" json:"-"`
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
	SubItems    []SubItem         `db:"-" json:"group_items,omitempty"`
}

// SubItem represents a product variant or grouped item.
type SubItem struct {
	ID           int64             `db:"id" json:"id,omitempty"`
	ProductID    int64             `db:"product_id" json:"-"`
	Title        string            `db:"title" json:"title"`
	Subtitle     *string           `db:"subtitle" json:"subtitle,omitempty"`
	Brand        *string           `db:"brand" json:"brand,omitempty"`
	Img          *string           `db:"img" json:"img,omitempty"`
	Category     []string          `db:"category" json:"category,omitempty"`
	Description  *string           `db:"description" json:"description,omitempty"`
	Code         *string           `db:"code" json:"code,omitempty"`
	Links        map[string]string `db:"-" json:"links,omitempty"`
	ShopeeLink   *string           `db:"shopee_link" json:"-"`
	TiktokLink   *string           `db:"tiktok_link" json:"-"`
	LazadaLink   *string           `db:"lazada_link" json:"-"`
	OtherLink    *string           `db:"other_link" json:"-"`
	DisplayOrder int16             `db:"display_order" json:"display_order,omitempty"`
	CreatedAt    time.Time         `db:"created_at" json:"-"`
	UpdatedAt    time.Time         `db:"updated_at" json:"-"`
}

// ToLinks converts individual link fields to Links map
func (s *SubItem) ToLinks() {
	s.Links = map[string]string{}
	if s.ShopeeLink != nil {
		s.Links["shopee"] = *s.ShopeeLink
	}
	if s.TiktokLink != nil {
		s.Links["tiktok"] = *s.TiktokLink
	}
	if s.LazadaLink != nil {
		s.Links["lazada"] = *s.LazadaLink
	}
	if s.OtherLink != nil {
		s.Links["other"] = *s.OtherLink
	}
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
