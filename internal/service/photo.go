package service

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
)

// PhotoService provides gallery photos
type PhotoService interface {
	Photos(ctx context.Context) ([]gallery.Photo, error)
}
