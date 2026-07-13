package gallery

import (
	"context"

	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
)

// Builder constructs gallery photos
type Builder struct {
	reader *metadata.Reader
}

func New(reader *metadata.Reader) *Builder {
	return &Builder{
		reader: reader,
	}
}

func (b *Builder) Build(
	ctx context.Context,
	files []scanner.File,
) ([]Photo, error) {
	return nil, nil
}
