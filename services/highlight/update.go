package highlight

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (s *HighlightService) Update(ctx context.Context, id int64, h domain.Highlight) (*domain.Highlight, bool, error) {
	return s.highlightRepo.Update(ctx, id, h)
}
