package gallery

import (
	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
)

// Photo represents an image in the gallery.
//
// A Photo combines filesystem information with extracted metadata
// and a stable identifier.
type Photo struct {
	ID string

	scanner.File
	metadata.Metadata
}
