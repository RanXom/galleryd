package service

import (
	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/scanner"
)

func NewGalleryService(
	scanner *scanner.Scanner,
	builder *gallery.Builder,
) *GalleryService {
	return &GalleryService{
		scanner: scanner,
		builder: builder,
	}
}
