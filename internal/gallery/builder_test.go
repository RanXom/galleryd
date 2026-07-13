package gallery

import (
	"context"
	"testing"
	"time"

	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
	"github.com/RanXom/galleryd/internal/testfixtures"
)

func TestBuilder(t *testing.T) {
	reader := metadata.New()
	builder := New(reader)

	t.Run("builds gallery photos", func(t *testing.T) {
		files := []scanner.File{
			{
				Path:         testfixtures.Canon40D(),
				Root:         "testfixtures",
				RelativePath: "Canon_40D.jpg",
				ModTime:      time.Now(),
			},
		}

		photos, err := builder.Build(context.Background(), files)
		if err != nil {
			t.Fatalf("build gallery: %v", err)
		}

		if len(photos) != 1 {
			t.Fatalf("expected 1 photo, got %d", len(photos))
		}
	})

	t.Run("generates deterministic ids", func(t *testing.T) {
		file := scanner.File{
			Path:         testfixtures.Canon40D(),
			Root:         "photos",
			RelativePath: "Canon_40D.jpg",
			ModTime:      time.Now(),
		}

		first, err := builder.Build(context.Background(), []scanner.File{file})
		if err != nil {
			t.Fatal(err)
		}

		second, err := builder.Build(context.Background(), []scanner.File{file})
		if err != nil {
			t.Fatal(err)
		}

		if first[0].ID != second[0].ID {
			t.Fatal("photo IDs should be deterministic")
		}
	})

	t.Run("respects context cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := builder.Build(ctx, nil)
		if err == nil {
			t.Fatal("expected context cancellation error")
		}
	})
}
