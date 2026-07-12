package scanner

import (
	"context"
	"path/filepath"
)

// Scanner discovers image files from one or more filesystem roots
//
// Scanner is safe for concurrent use after construction,
// provided Scan is not called concurrently
type Scanner struct {
	config Config
}

// New creates a Scanner using the supplied configuration.
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

// Scan recursively walks every configured root and returns the
// discovered image files.
//
// Scan respects context cancellation.
//
// Returned files are deduplicated by absolute path.
//
// Results are not sorted.
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
