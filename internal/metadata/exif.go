package metadata

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// readEXIF decodes the EXIF metadata from an image.
//
// Only the first 64 KiB are read since EXIF metadata is stored
// near the beginning of JPEG files.
func readEXIF(path string) (*exif.Exif, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open image: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	reader := io.LimitReader(file, 64*1024)

	x, err := exif.Decode(reader)
	if err != nil {
		return nil, err
	}

	return x, nil
}

// dateTakenFromEXIF extracts the DateTime EXIF tag.
func dateTakenFromEXIF(x *exif.Exif) (time.Time, error) {
	return x.DateTime()
}

// orientationFromEXIF extracts the EXIF Orientation tag.
//
// If the image has no orientation tag, the default orientation (1)
// is returned.
func orientationFromEXIF(x *exif.Exif) (int, error) {
	tag, err := x.Get(exif.Orientation)
	if err != nil {
		return 1, nil
	}

	orientation, err := tag.Int(0)
	if err != nil {
		return 1, err
	}

	return orientation, nil
}
