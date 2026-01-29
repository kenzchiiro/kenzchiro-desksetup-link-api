package highlight

import (
	"context"
	"encoding/json"
	"time"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (r *HighlightRepository) List(ctx context.Context) ([]domain.Highlight, error) {
	ctx, cancel := context.WithTimeout(ctx, r.timeout)
	defer cancel()

	type highlightWithProduct struct {
		domain.Highlight
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
