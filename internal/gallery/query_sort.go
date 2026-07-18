package gallery

import "sort"

func sortQueryPhotos(
	photos []Photo,
	query Query,
) {
	if query.Sort == "" {
		return
	}

	sort.SliceStable(photos, func(i, j int) bool {
		switch query.Sort {
		case SortByPath:
			if query.Order == SortDesc {
				return photos[i].RelativePath > photos[j].RelativePath
			}

			return photos[i].RelativePath < photos[j].RelativePath
		case SortByDateTime:
			if query.Order == SortAsc {
				if !photos[i].DateTaken.Equal(photos[j].DateTaken) {
					return photos[i].DateTaken.Before(photos[j].DateTaken)
				}

				return photos[i].RelativePath < photos[j].RelativePath
			}

			if !photos[i].DateTaken.Equal(photos[j].DateTaken) {
				return photos[i].DateTaken.After(photos[j].DateTaken)
			}

			return photos[i].RelativePath < photos[j].RelativePath

		default:
			return false
		}
	})
}
