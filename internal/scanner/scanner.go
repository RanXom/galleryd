package scanner

import (
	"context"
	"path/filepath"
)

type Scanner struct {
	roots []string
}

func New(roots []string) *Scanner {
	clean := make([]string, len(roots))
	for i, root := range roots {
		clean[i] = filepath.Clean(root)
	}

	return &Scanner{
		roots: clean,
	}
}

// Scan walks every configured root and returns the discovered files.
//
// The returned slice is not guaranteed to be sorted.
func (s *Scanner) Scan(ctx context.Context) ([]File, error) {
	files := make([]File, 0)
	seen := make(map[string]struct{})

	for _, root := range s.roots {
		discovered, err := s.walk(ctx, root)
		if err != nil {
			return nil, err
		}

		for _, file := range discovered {
			if _, ok := seen[file.Path]; ok {
				continue
			}

			seen[file.Path] = struct{}{}
			files = append(files, file)
		}
	}

	return files, nil
}
