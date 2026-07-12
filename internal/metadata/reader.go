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
		Width:     width,
		Height:    height,
		DateTaken: file.ModTime,
	}

	dateTaken, err := readDateTaken(file.Path)
	if err == nil {
		metadata.DateTaken = dateTaken
	}

	orientation, err := readOrientation(file.Path)
	if err == nil {
		metadata.Orientation = orientation
	} else {
		metadata.Orientation = 1
	}

	return metadata, nil
}
