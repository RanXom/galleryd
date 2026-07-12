package scanner

import (
	"context"
	"io/fs"
	"path/filepath"
)

func (s *Scanner) walk(ctx context.Context, root string) ([]File, error) {
	files := make([]File, 0)

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err != nil {
			return err
		}

		if d.IsDir() {
			if shouldSkipDir(d.Name()) {
				return filepath.SkipDir
			}

			return nil
		}

		if !isImage(path) {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		files = append(files, File{
			Path:         path,
			Root:         root,
			RelativePath: rel,
			Size:         info.Size(),
			ModTime:      info.ModTime(),
		})

		return nil
	})

	return files, err
}
