package scanner

import (
	"context"
	"path/filepath"
)

type Scanner struct {
	config Config
}

func New(config Config) *Scanner {
	clean := make([]string, len(config.Roots))
	for i, root := range config.Roots {
		clean[i] = filepath.Clean(root)
	}

	config.Roots = clean
	return &Scanner{
		config: config,
	}
}

// Scan walks every configured root and returns the discovered files.
//
// The returned slice is not guaranteed to be sorted.
func (s *Scanner) Scan(ctx context.Context) ([]File, error) {
	files := make([]File, 0)
	seen := make(map[string]struct{})

	for _, root := range s.config.Roots {
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
