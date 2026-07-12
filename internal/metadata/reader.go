package metadata

import (
	"context"

	"github.com/RanXom/galleryd/internal/scanner"
)

type Reader struct{}

func New() *Reader {
	return &Reader{}
}

func (r *Reader) Read(ctx context.Context, file scanner.File) (Metadata, error) {
	return Metadata{}, nil
}
