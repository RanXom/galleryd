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

// readOrientation extracts the EXIF Orientation tag.
//
// If the image has no orientation tag, the default orientation (1)
// is returned
func readOrientation(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 1, fmt.Errorf("open image: %w", err)
	}
	defer func() {
		_ = file.Close()
	}()

	reader := io.LimitReader(file, 64*1024)

	x, err := exif.Decode(reader)
	if err != nil {
		return 1, err
	}

	tag, err := x.Get(exif.Orientation)
	if err != nil {
		return 1, err
	}

	orientaion, err := tag.Int(0)
	if err != nil {
		return 1, err
	}

	return orientaion, nil
}
