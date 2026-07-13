package thumbnail

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// initializeCache prepares the thumbnail cache directory.
//
// It creates the current cache generation and removes obsolete cache
// generations left behind by previous versions.
func (g *Generator) initializeCache() error {
	thumbsDir := filepath.Join(g.cacheDir, "thumbs")

	if err := os.MkdirAll(
		filepath.Join(thumbsDir, cacheVersion()),
		0o775,
	); err != nil {
		return fmt.Errorf("create cache directory: %w", err)
	}

	entries, err := os.ReadDir(thumbsDir)
	if err != nil {
		return fmt.Errorf("read cache directory: %w", err)
	}

	current := cacheVersion()

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name()

		if !strings.HasPrefix(name, "v") {
			continue
		}

		if name == current {
			continue
		}

		if err := os.RemoveAll(filepath.Join(thumbsDir, name)); err != nil {
			return fmt.Errorf(
				"remove obsolete cache %q: %w",
				name,
				err,
			)
		}
	}

	return nil
}
