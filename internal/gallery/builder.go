package gallery

import (
	"context"

	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
)

// Builder constructs gallery photos from scanned files.
type Builder struct {
	reader *metadata.Reader
}

func New(reader *metadata.Reader) *Builder {
	return &Builder{
		reader: reader,
	}
}

// Build converts scanned files into gallery photos.
//
// Photos are sorted by capture date in descending order.
func (b *Builder) Build(
	ctx context.Context,
	files []scanner.File,
) ([]Photo, error) {
	photos := make([]Photo, 0, len(files))

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

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
