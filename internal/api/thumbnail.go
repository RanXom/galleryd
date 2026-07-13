package api

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/thumbnail"
)

// ThumbnailGenerator generates thumbnails.
type ThumbnailGenerator interface {
	Generate(
		ctx context.Context,
		photo gallery.Photo,
	) (thumbnail.Thumbnail, error)
}
