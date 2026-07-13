package service

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
)

// Gallery scans the configured roots and builds the gallery.
func (s *galleryService) Gallery(ctx context.Context) ([]gallery.Photo, error) {
	files, err := s.scanner.Scan(ctx)
	if err != nil {
		return nil, err
	}

	return s.builder.Build(ctx, files)
}
