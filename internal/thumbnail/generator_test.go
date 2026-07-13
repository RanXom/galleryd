package thumbnail

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/scanner"
	"github.com/RanXom/galleryd/internal/testfixtures"
)

func TestGenerator(t *testing.T) {
	cacheDir := t.TempDir()

	generator, err := New(cacheDir)
	if err != nil {
		t.Fatal(err)
	}

	photo := gallery.Photo{
		ID: "canon40d",

		File: scanner.File{
			Path:         testfixtures.Canon40D(),
			Root:         "photos",
			RelativePath: "Canon_40D.jpg",
			ModTime:      time.Now(),
		},
	}

	t.Run("generates thumbnail", func(t *testing.T) {
		thumb, err := generator.Generate(
			context.Background(),
			photo,
		)
		if err != nil {
			t.Fatal(err)
		}

		if _, err := os.Stat(thumb.Path); err != nil {
			t.Fatalf("thumbnail was not created: %v", err)
		}
	})

	t.Run("reuses cached thumbnail", func(t *testing.T) {
		first, err := generator.Generate(
			context.Background(),
			photo,
		)
		if err != nil {
			t.Fatal(err)
		}

		info1, err := os.Stat(first.Path)
		if err != nil {
			t.Fatal(err)
		}

		time.Sleep(10 * time.Millisecond)

		second, err := generator.Generate(
			context.Background(),
			photo,
		)
		if err != nil {
			t.Fatal(err)
		}

		info2, err := os.Stat(second.Path)
		if err != nil {
			t.Fatal(err)
		}

		if !info1.ModTime().Equal(info2.ModTime()) {
			t.Fatal("thumbnail was regenerated instead of reused")
		}
	})

	t.Run("respects cancelled context", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := generator.Generate(ctx, photo)

		if err == nil {
			t.Fatal("expected context cancellation")
		}
	})
}
