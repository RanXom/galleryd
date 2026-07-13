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

	photos, err := s.builder.Build(ctx, files)
	if err != nil {
		return nil, err
	}

	index := galleryIndex{
		photos: photos,
		byID:   make(map[string]gallery.Photo, len(photos)),
	}

	for _, photo := range photos {
		index.byID[photo.ID] = photo
	}

	s.index = index

	return s.index.photos, nil
}
