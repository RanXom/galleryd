package service

import (
	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/scanner"
)

// GalleryService builds gallery photos
type GalleryService struct {
	scanner *scanner.Scanner
	builder *gallery.Builder
}
