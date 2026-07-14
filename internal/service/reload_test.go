package service

import (
	"context"
	"testing"

	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
)

func TestReloadInitialLoad(t *testing.T) {
	ctx := context.Background()

	sc := scanner.New(scanner.Config{
		Roots: []string{
			"../scanner/testdata/gallery",
		},
	})

	builder := gallery.New(metadata.New())

	service := New(sc, builder)

	if err := service.Reload(ctx); err != nil {
		t.Fatalf("reload gallery: %v", err)
	}

	photos, err := service.Gallery(ctx)
	if err != nil {
		t.Fatalf("gallery: %v", err)
	}

	if len(photos) == 0 {
		t.Fatal("expected gallery to contain photos")
	}

	for _, photo := range photos {
		got, err := service.Photo(ctx, photo.ID)
		if err != nil {
			t.Fatalf("lookup photo %q: %v", photo.ID, err)
		}

		if got.ID != photo.ID {
			t.Fatalf("expected id %q, got %q", photo.ID, got.ID)
		}
	}
}
