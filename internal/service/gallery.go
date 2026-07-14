package service

import (
	"sync"

	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/scanner"
)

// galleryService implements GalleryService.
type galleryService struct {
	scanner *scanner.Scanner
	builder *gallery.Builder

	mu sync.RWMutex

	index galleryIndex
}
