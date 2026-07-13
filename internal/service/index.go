package service

import "github.com/RanXom/galleryd/internal/gallery"

// galleryIndex stores an in-memory snapshot of the gallery.
type galleryIndex struct {
	photos []gallery.Photo
	byID   map[string]gallery.Photo
}
