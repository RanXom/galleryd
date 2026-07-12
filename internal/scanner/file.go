package scanner

import "time"

// File represents a discovered file on disk.
//
// It contains only filesystem information.
// Higher-level packages enrich it with metadata such as EXIF,
// dimensions, IDs and thumbnails.
type File struct {
	Path         string
	Root         string
	RelativePath string

	Size    int64
	ModTime time.Time
}
