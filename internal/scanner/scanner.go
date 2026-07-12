package scanner

import "context"

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
	var files []File

	for _, root := range s.roots {
		discovered, err := s.walk(ctx, root)
		if err != nil {
			return nil, err
		}

		files = append(files, discovered...)
	}

	return nil, nil
}
