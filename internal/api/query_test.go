package api

import (
	"net/http/httptest"
	"testing"

	"github.com/RanXom/galleryd/internal/gallery"
)

func TestParseQueryDefaults(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos",
		nil,
	)

	query, err := parseQuery(req)
	if err != nil {
		t.Fatalf("parse query: %v", err)
	}

	if query.Limit != 0 {
		t.Fatalf(
			"expected limit %d, got %d",
			0,
			query.Limit,
		)
	}

	if query.Offset != 0 {
		t.Fatalf(
			"expected offset %d, got %d",
			0,
			query.Offset,
		)
	}
}

func TestParseQueryPagination(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?limit=25&offset=100",
		nil,
	)

	query, err := parseQuery(req)
	if err != nil {
		t.Fatalf("parse query: %v", err)
	}

	if query.Limit != 25 {
		t.Fatalf(
			"expected limit %d, got %d",
			25,
			query.Limit,
		)
	}

	if query.Offset != 100 {
		t.Fatalf(
			"expected offset %d, got %d",
			100,
			query.Offset,
		)
	}
}

func TestParseQueryInvalidLimit(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?limit=abc",
		nil,
	)

	if _, err := parseQuery(req); err == nil {
		t.Fatal("expected error")
	}
}

func TestParseQueryInvalidOffset(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?offset=abc",
		nil,
	)

	if _, err := parseQuery(req); err == nil {
		t.Fatal("expected error")
	}
}

func TestParseQueryRejectsNegativeValues(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?limit=-1&offset=-5",
		nil,
	)

	if _, err := parseQuery(req); err == nil {
		t.Fatal("expected error")
	}
}

func TestParseQuerySorting(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?sort=path&order=asc",
		nil,
	)

	query, err := parseQuery(req)
	if err != nil {
		t.Fatalf("parse query: %v", err)
	}

	if query.Sort != gallery.SortByPath {
		t.Fatalf(
			"expected sort %q, got %q",
			gallery.SortByPath,
			query.Sort,
		)
	}

	if query.Order != gallery.SortAsc {
		t.Fatalf(
			"expected order %q, got %q",
			gallery.SortAsc,
			query.Order,
		)
	}
}

func TestParseQueryInvalidSort(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?sort=banana",
		nil,
	)

	if _, err := parseQuery(req); err == nil {
		t.Fatal("expected error")
	}
}

func TestParseQueryInvalidOrder(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?order=sideways",
		nil,
	)

	if _, err := parseQuery(req); err == nil {
		t.Fatal("expected error")
	}
}

func TestParseQuerySortWithoutOrder(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?sort=path",
		nil,
	)

	query, err := parseQuery(req)
	if err != nil {
		t.Fatalf("parse query: %v", err)
	}

	if query.Sort != gallery.SortByPath {
		t.Fatalf(
			"expected sort %q, got %q",
			gallery.SortByPath,
			query.Sort,
		)
	}

	if query.Order != "" {
		t.Fatalf(
			"expected empty order, got %q",
			query.Order,
		)
	}
}

func TestParseQueryOrderWithoutSort(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"/api/photos?order=desc",
		nil,
	)

	query, err := parseQuery(req)
	if err != nil {
		t.Fatalf("parse query: %v", err)
	}

	if query.Sort != "" {
		t.Fatalf(
			"expected empty sort, got %q",
			query.Sort,
		)
	}

	if query.Order != gallery.SortDesc {
		t.Fatalf(
			"expected order %q, got %q",
			gallery.SortDesc,
			query.Order,
		)
	}
}
