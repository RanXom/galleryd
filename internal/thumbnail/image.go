package thumbnail

import (
	"fmt"
	"image"

	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/disintegration/imaging"
)

const thumbnailSize = 512

// generateImage decodes and resizes a photo.
//
// The returned image is not written to disk.
func generateImage(photo gallery.Photo) (image.Image, error) {
	src, err := imaging.Open(photo.Path)
	if err != nil {
		return nil, fmt.Errorf("open image: %w", err)
	}

	thumb := imaging.Fit(
		src,
		thumbnailSize,
		thumbnailSize,
		imaging.Lanczos,
	)

	return thumb, nil
}
