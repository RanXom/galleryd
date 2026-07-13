package thumbnail

import (
	"context"

	"github.com/RanXom/galleryd/internal/gallery"
)

// Generator creates thumbnails for gallery photos
type Generator struct {
	cacheDir string
}

func New(cacheDir string) (*Generator, error) {
	g := &Generator{
		cacheDir: cacheDir,
	}

	if err := g.initializeCache(); err != nil {
		return nil, err
	}

	return g, nil
}

// Generate creates (or retrieves) a thumbnail
func (g *Generator) Generate(
	ctx context.Context,
	photo gallery.Photo,
) (Thumbnail, error) {
	select {
	case <-ctx.Done():
		return Thumbnail{}, ctx.Err()
	default:
	}

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
