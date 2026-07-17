package api

import (
	"net/http/httptest"
	"testing"
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
