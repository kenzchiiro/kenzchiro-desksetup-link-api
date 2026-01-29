package highlight

import (
	"context"
)

func (s *HighlightService) Delete(ctx context.Context, id int64) (bool, error) {
	return s.highlightRepo.Delete(ctx, id)
}
