package scanner

import (
	"context"
	"path/filepath"
	"testing"
)

func TestScanner(t *testing.T) {
	root := filepath.Join("testdata", "gallery")

	s := New(Config{
		Roots: []string{root},
	})

	files, err := s.Scan(context.Background())
	if err != nil {
		t.Fatalf("scan failed: %v", err)
	}

	t.Run("find supported images", func(t *testing.T) {
		expected := map[string]struct{}{
			"cat.jpg":    {},
			"dog.png":    {},
			"beach.webp": {},
		}

		for _, file := range files {
			delete(expected, filepath.Base(file.Path))
		}

		if len(expected) != 0 {
			t.Fatalf("missing files: %v", expected)
		}
	})

	t.Run("ignore unsupported files", func(t *testing.T) {
		for _, file := range files {
			switch filepath.Base(file.Path) {
			case "readme.txt", "jeffsfiles.pdf":
				t.Fatalf("scanner returned unsupported files %q", file.Path)
			}
		}
	})

	t.Run("skip hidden directories", func(t *testing.T) {
		for _, file := range files {
			if filepath.Base(file.Path) == "nuclearlaunchcode.txt" {
				t.Fatal("hidden file entered the scanner")
			}
		}
	})

	t.Run("deduplicates overlapping roots", func(t *testing.T) {
		root := filepath.Join("testdata", "duplicate")

		s := New(Config{
			Roots: []string{
				root,
				root,
			},
		})

		files, err := s.Scan(context.Background())
		if err != nil {
			t.Fatalf("scan failed: %v", err)
		}

		if len(files) != 2 {
			t.Fatalf("expected 2 unique files, got %d", len(files))
		}
	})

	t.Run("ignores symbolic links", func(t *testing.T) {
		for _, file := range files {
			if filepath.Base(file.Path) == "link.jpg" {
				t.Fatal("scanner followed symbolic link")
			}
		}
	})
}
