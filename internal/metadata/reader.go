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
	select {
	case <-ctx.Done():
		return Metadata{}, ctx.Err()
	default:
	}

	width, height, err := readDimensions(file.Path)
	if err != nil {
		return Metadata{}, err
	}

	metadata := Metadata{
		Width:  width,
		Height: height,
	}

	dateTaken, err := readDateTaken(file.Path)
	if err == nil {
		metadata.DateTaken = dateTaken
	}

	return metadata, nil
}
