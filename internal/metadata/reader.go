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

	metadata.Orientation = 1

	x, err := readEXIF(file.Path)
	if err == nil {
		if dateTaken, err := dateTakenFromEXIF(x); err == nil {
			metadata.DateTaken = dateTaken
		}

		if orientation, err := orientationFromEXIF(x); err == nil {
			metadata.Orientation = orientation
		}
	}

	return metadata, nil
}
