package thumbnail

import (
	"errors"
	"os"
)

// thumbnailExists reports wether a cached thumbnail exists
func thumbnailExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, err
}
