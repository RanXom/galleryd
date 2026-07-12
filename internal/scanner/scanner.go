package scanner

import "context"

type Scanner struct {
	roots []string
}

func New(roots []string) *Scanner {
	return &Scanner{
		roots: roots,
	}
}

// Scan walks every configured root and returns the discovered files.
//
// The returned slice is not guaranteed to be sorted.
func (s *Scanner) Scan(ctx context.Context) ([]File, error) {
	return nil, nil
}
