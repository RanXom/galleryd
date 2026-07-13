package service

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
)

// Load scans the configured roots and rebuilds the in-memory gallery.
func (s *galleryService) Load(ctx context.Context) error {
	files, err := s.scanner.Scan(ctx)
	if err != nil {
		return err
	}

	photos, err := s.builder.Build(ctx, files)
	if err != nil {
		return err
	}

	index := galleryIndex{
		photos: photos,
		byID:   make(map[string]gallery.Photo, len(photos)),
	}

	for _, photo := range photos {
		index.byID[photo.ID] = photo
	}

	s.index = index

	return nil
}
