package gallery

import (
	"path/filepath"
	"strings"
)

func filterQueryPhotos(
	photos []Photo,
	query Query,
) []Photo {
	filtered := make([]Photo, 0, len(photos))

	for _, photo := range photos {
		if !matchesExtension(photo, query.Extension) {
			continue
		}

		if !matchesSearch(photo, query.Search) {
			continue
		}

		filtered = append(filtered, photo)
	}

	return filtered
}

func matchesExtension(
	photo Photo,
	extension string,
) bool {
	if extension == "" {
		return true
	}

	ext := strings.TrimPrefix(
		filepath.Ext(photo.RelativePath),
		".",
	)

	extension = strings.TrimPrefix(
		extension,
		".",
	)

	return strings.EqualFold(ext, extension)
}

func matchesSearch(
	photo Photo,
	search string,
) bool {
	if search == "" {
		return true
	}

	return strings.Contains(
		strings.ToLower(photo.RelativePath),
		strings.ToLower(search),
	)
}
