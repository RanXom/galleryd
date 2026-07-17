package service

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
	"github.com/RanXom/galleryd/internal/testfixtures"
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

	photos, err := service.Gallery(ctx, gallery.Query{})
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

func TestReloadReplacesGallery(t *testing.T) {
	ctx := context.Background()

	root := t.TempDir()

	copyFile := func(src, dst string) {
		t.Helper()

		data, err := os.ReadFile(src)
		if err != nil {
			t.Fatalf("read fixture: %v", err)
		}

		if err := os.WriteFile(dst, data, 0o644); err != nil {
			t.Fatalf("write fixture: %v", err)
		}
	}

	first := filepath.Join(root, "first.jpg")
	second := filepath.Join(root, "second.jpg")

	copyFile(testfixtures.Canon40D(), first)

	sc := scanner.New(scanner.Config{
		Roots: []string{root},
	})

	builder := gallery.New(metadata.New())

	service := New(sc, builder)

	if err := service.Reload(ctx); err != nil {
		t.Fatalf("first reload: %v", err)
	}

	before, err := service.Gallery(ctx, gallery.Query{})
	if err != nil {
		t.Fatalf("gallery: %v", err)
	}

	if len(before) != 1 {
		t.Fatalf("expected 1 photo, got %d", len(before))
	}

	oldID := before[0].ID

	if err := os.Remove(first); err != nil {
		t.Fatalf("remove first image: %v", err)
	}

	copyFile(testfixtures.NikonD70(), second)

	if err := service.Reload(ctx); err != nil {
		t.Fatalf("second reload: %v", err)
	}

	after, err := service.Gallery(ctx, gallery.Query{})
	if err != nil {
		t.Fatalf("gallery: %v", err)
	}

	if len(after) != 1 {
		t.Fatalf("expected 1 photo, got %d", len(after))
	}

	if after[0].ID == oldID {
		t.Fatal("gallery snapshot was not replaced")
	}

	if _, err := service.Photo(ctx, oldID); err != ErrPhotoNotFound {
		t.Fatalf("expected ErrPhotoNotFound, got %v", err)
	}
}
