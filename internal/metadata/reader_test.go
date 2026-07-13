package metadata

import (
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/RanXom/galleryd/internal/scanner"
)

func TestReader(t *testing.T) {
	reader := New()

	t.Run("reads image dimensions", func(t *testing.T) {
		file := scanner.File{
			Path: filepath.Join("testdata", "Canon_40D.jpg"),
		}

		metadata, err := reader.Read(context.Background(), file)
		if err != nil {
			t.Fatalf("read metadata: %v", err)
		}

		if metadata.Width == 0 || metadata.Height == 0 {
			t.Fatal("expected image dimensions")
		}
	})

	t.Run("reads EXIF timestamp", func(t *testing.T) {
		file := scanner.File{
			Path: filepath.Join("testdata", "Canon_40D.jpg"),
		}

		metadata, err := reader.Read(context.Background(), file)
		if err != nil {
			t.Fatalf("read metadata: %v", err)
		}

		expected := time.Date(2008, time.May, 30, 15, 56, 1, 0, time.Local)

		if !metadata.DateTaken.Equal(expected) {
			t.Fatalf("expected %v, got %v", expected, metadata.DateTaken)
		}
	})

	t.Run("falls back to modification time", func(t *testing.T) {
		modTime := time.Date(
			2025,
			time.January,
			1,
			12,
			0,
			0,
			0,
			time.UTC,
		)

		file := scanner.File{
			Path:    filepath.Join("testdata", "no-exif.png"),
			ModTime: modTime,
		}

		metadata, err := reader.Read(context.Background(), file)
		if err != nil {
			t.Fatalf("read metadata: %v", err)
		}

		if !metadata.DateTaken.Equal(modTime) {
			t.Fatalf(
				"expected fallback modification time %v, got %v",
				modTime,
				metadata.DateTaken,
			)
		}
	})

	t.Run("returns default orientation", func(t *testing.T) {
		file := scanner.File{
			Path: filepath.Join("testdata", "no-exif.png"),
		}

		metadata, err := reader.Read(context.Background(), file)
		if err != nil {
			t.Fatalf("read metadata: %v", err)
		}

		if metadata.Orientation != 1 {
			t.Fatalf(
				"expected orientation 1, got %d",
				metadata.Orientation,
			)
		}
	})
}
