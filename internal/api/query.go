package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RanXom/galleryd/internal/gallery"
)

func parsePositiveInt(value string) (int, error) {
	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, err
	}

	if n < 0 {
		return 0, errors.New("must be non-negative")
	}

	return n, nil
}

func parseQuery(r *http.Request) (gallery.Query, error) {
	query := gallery.Query{}

	values := r.URL.Query()

	if limit := values.Get("limit"); limit != "" {
		n, err := parsePositiveInt(limit)
		if err != nil {
			return gallery.Query{}, err
		}

		query.Limit = n
	}

	if offset := values.Get("offset"); offset != "" {
		n, err := parsePositiveInt(offset)
		if err != nil {
			return gallery.Query{}, err
		}

		query.Offset = n
	}

	if sort := values.Get("sort"); sort != "" {
		switch gallery.SortField(sort) {
		case gallery.SortByDateTime, gallery.SortByPath:
			query.Sort = gallery.SortField(sort)
		default:
			return gallery.Query{}, errors.New("invalid sort field")
		}
	}

	if order := values.Get("order"); order != "" {
		switch gallery.SortOrder(order) {
		case gallery.SortAsc, gallery.SortDesc:
			query.Order = gallery.SortOrder(order)
		default:
			return gallery.Query{}, errors.New("invalid sort order")
		}
	}

	return query, nil
}
