package api

func thumbnailURL(id string) string {
	return "/thumb/" + id
}

func photoURL(id string) string {
	return "/photo/" + id
}
