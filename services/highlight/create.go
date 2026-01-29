package highlight

import (
	"context"

	"github.com/kenzchiro/desksetup-link-api/domain"
)

func (s *HighlightService) Create(ctx context.Context, h domain.Highlight) (domain.Highlight, error) {
	return s.highlightRepo.Create(ctx, h)
}
