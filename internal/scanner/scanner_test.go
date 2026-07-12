package scanner

import (
	"context"
	"path/filepath"
	"testing"
)

func TestScan(t *testing.T) {
	root := filepath.Join("testdata", "gallery")

	s := New([]string{root})

	files, err := s.Scan(context.Background())
	if err != nil {
		t.Fatalf("Scan failed: %v", err)
	}

	t.Run("Find supported images", func(t *testing.T) {
		expected := map[string]struct{}{
			"cat.jpg":    {},
			"dog.png":    {},
			"beach.webp": {},
		}

		for _, file := range files {
			delete(expected, filepath.Base(file.Path))
		}

		if len(expected) != 0 {
			t.Fatalf("Missing files: %v", expected)
		}
	})

	t.Run("Ignore unsupported files", func(t *testing.T) {
		for _, file := range files {
			switch filepath.Base(file.Path) {
			case "readme.txt", "jeffsfiles.pdf":
				t.Fatalf("Scanner returned unsupported files %q", file.Path)
			}
		}
	})

	t.Run("Skip hidden directories", func(t *testing.T) {
		for _, file := range files {
			if filepath.Base(file.Path) == "nuclearlaunchcode.txt" {
				t.Fatal("Hidden file entered the scanner")
			}
		}
	})

	t.Run("Deduplicates overlapping roots", func(t *testing.T) {
		root := filepath.Join("testdata", "duplicate")

		s := New([]string{
			root,
			root,
		})

		files, err := s.Scan(context.Background())
		if err != nil {
			t.Fatalf("Scan failed: %v", err)
		}

		if len(files) != 2 {
			t.Fatalf("Expected 2 unique files, got %d", len(files))
		}
	})
}
