package metadata

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rwcarlsen/goexif/exif"
)

// readDateTaken extracts the DateTime EXIF tag from an image
//
// Only the first 64 KiB of the file are read since EXIF metadata
// is stored near the beginning of JPEG files.
func readDateTaken(path string) (time.Time, error) {
	file, err := os.Open(path)
	if err != nil {
		return time.Time{}, fmt.Errorf("open image: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	reader := io.LimitReader(file, 64*1024)

	x, err := exif.Decode(reader)
	if err != nil {
		return time.Time{}, err
	}

	dateTaken, err := x.DateTime()
	if err != nil {
		return time.Time{}, err
	}

	return dateTaken, nil
}
