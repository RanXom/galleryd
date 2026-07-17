package gallery

func QueryPhotos(
	photos []Photo,
	query Query,
) []Photo {
	if query.Offset >= len(photos) {
		return []Photo{}
	}

	photos = photos[query.Offset:]

	if query.Limit == 0 || query.Limit >= len(photos) {
		return photos
	}

	return photos[:query.Limit]
}
