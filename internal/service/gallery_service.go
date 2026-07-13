package service

import (
	"context"
	"errors"

	"github.com/RanXom/galleryd/internal/gallery"
)

// Gallery scans the configured roots and builds the gallery.
func (s *galleryService) Gallery(ctx context.Context) ([]gallery.Photo, error) {
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
