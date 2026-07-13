package thumbnail

import (
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/chai2010/webp"
)

// saveImage writes a thumbnail image to disk as WebP.
func saveImage(path string, img image.Image) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("created cache directory: %w", err)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create thumbnail: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	options := &webp.Options{
		Lossless: false,
		Quality:  80,
	}

	if err := webp.Encode(file, img, options); err != nil {
		return fmt.Errorf("encode webp: %w", err)
	}

	return nil
}
