package service

import (
	"context"
	"errors"

	"github.com/RanXom/galleryd/internal/gallery"
)

// Gallery scans the configured roots and builds the gallery.
func (s *galleryService) Gallery(ctx context.Context, query gallery.Query) ([]gallery.Photo, error) {
	_ = query

	s.mu.RLock()
	defer s.mu.RUnlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	if s.index.byID == nil {
		return nil, errors.New("gallery not loaded")
	}

	return s.index.photos, nil
}
