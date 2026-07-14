package service

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
)

// GalleryService provides gallery operations.
type GalleryService interface {
	// Load scans the configured roots and builds the in-memory gallery.
	Reload(ctx context.Context) error

	// Gallery returns the currently loaded gallery.
	Gallery(ctx context.Context) ([]gallery.Photo, error)

	// Photo returns a single photo by its ID.
	Photo(ctx context.Context, id string) (gallery.Photo, error)
}
