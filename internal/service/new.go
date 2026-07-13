package service

import (
	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/scanner"
)

func New(
	scanner *scanner.Scanner,
	builder *gallery.Builder,
) GalleryService {
	return &galleryService{
		scanner: scanner,
		builder: builder,
	}
}
