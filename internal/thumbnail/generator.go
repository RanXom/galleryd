package thumbnail

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
)

// Generator creates thumbnails for gallery photos
type Generator struct {
	cacheDir string
}

func New(cacheDir string) *Generator {
	return &Generator{
		cacheDir: cacheDir,
	}
}

// Generate creates (or retrieves) a thumbnail
func (g *Generator) Generate(
	ctx context.Context,
	photo gallery.Photo,
) (Thumbnail, error) {
	path := g.cachePath(photo)

	exists, err := thumbnailExists(path)
	if err != nil {
		return Thumbnail{}, err
	}

	if !exists {
		img, err := generateImage(photo)
		if err != nil {
			return Thumbnail{}, err
		}

		if err := saveImage(path, img); err != nil {
			return Thumbnail{}, err
		}
	}

	return Thumbnail{
		Path: path,
	}, nil
}
