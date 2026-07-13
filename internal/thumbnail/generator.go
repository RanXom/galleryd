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
	_, err := generateImage(photo)
	if err != nil {
		return Thumbnail{}, err
	}

	return Thumbnail{
		Path: g.cachePath(photo),
	}, nil
}
