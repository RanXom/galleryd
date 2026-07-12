package scanner

import "time"

// File represents a discovered file on disk.
//
// It contains only filesystem information.
// Metadata such as EXIF, thumbnails and IDs are added by higher
// level packages.
type File struct {
	Path         string
	Root         string
	RelativePath string

	Size    int64
	ModTime time.Time
}
