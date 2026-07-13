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
	photos := make([]Photo, 0, len(files))

	for _, file := range files {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		md, err := b.reader.Read(ctx, file)
		if err != nil {
			return nil, err
		}

		photos = append(photos, Photo{
			ID:       generateID(file),
			File:     file,
			Metadata: md,
		})
	}

	sortPhotos(photos)

	return photos, nil
}
