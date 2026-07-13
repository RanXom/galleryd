package gallery

import "sort"

// sortPhotos orders photos by date taken (newest first).
//
// Photos with identical timestamps are ordered by relative path to
// provide deterministic output.
func sortPhotos(photos []Photo) {
	sort.SliceStable(photos, func(i, j int) bool {
		if !photos[i].DateTaken.Equal(photos[j].DateTaken) {
			return photos[i].DateTaken.After(photos[j].DateTaken)
		}

		return photos[i].RelativePath < photos[j].RelativePath
	})
}
