package api

import (
	"net/http"
	"strconv"

	"github.com/RanXom/galleryd/internal/gallery"
)

func parseQuery(r *http.Request) (gallery.Query, error) {
	query := gallery.Query{}

	values := r.URL.Query()

	if limit := values.Get("limit"); limit != "" {
		n, err := strconv.Atoi(limit)
		if err != nil {
			return gallery.Query{}, err
		}

		query.Limit = n
	}

	if offset := values.Get("offset"); offset != "" {
		n, err := strconv.Atoi(offset)
		if err != nil {
			return gallery.Query{}, err
		}

		query.Offset = n
	}

	return query, nil
}
