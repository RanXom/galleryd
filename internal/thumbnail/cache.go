package thumbnail

import (
	"fmt"
	"path/filepath"

	"github.com/RanXom/galleryd/internal/gallery"
)

const (
	thumbnailCacheVersion = 1
	thumbnailExtension    = ".webp"
)

func cacheVersion() string {
	return fmt.Sprintf("v%d", thumbnailCacheVersion)
}

// cachePath returns the on-disk cache location for a photo thumbnail.
//
// Thumbnails are sharded by the first two characters of the photo ID to
// avoid placing large numbers of files in a single directory.
func (g *Generator) cachePath(photo gallery.Photo) string {
	shard := photo.ID[:2]

	filename := photo.ID + thumbnailExtension

	return filepath.Join(
		g.cacheDir,
		"thumbs",
		cacheVersion(),
		shard,
		filename,
	)
}
