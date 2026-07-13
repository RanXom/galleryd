package service

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
)

// GalleryService provides gallery operations.
type GalleryService interface {
	Gallery(ctx context.Context) ([]gallery.Photo, error)
}
