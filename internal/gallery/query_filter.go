package gallery

import (
	"path/filepath"
	"strings"
)

func filterQueryPhotos(
	photos []Photo,
	query Query,
) []Photo {
	if query.Extension == "" {
		return photos
	}

	filtered := make([]Photo, 0, len(photos))

	for _, photo := range photos {
		ext := strings.TrimPrefix(filepath.Ext(photo.RelativePath), ".")

		if strings.EqualFold(
			strings.TrimPrefix(ext, "."),
			strings.TrimPrefix(query.Extension, "."),
		) {
			filtered = append(filtered, photo)
		}
	}

	return filtered
}
