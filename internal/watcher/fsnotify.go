package watcher

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

// Run starts watching the configured roots.
//
// Filesystem events are currently logged for debugging.
// Future commits will debounce these events and invoke OnChange.
func (w *Watcher) Run(ctx context.Context) error {
	fswatcher, err := fsnotify.NewWatcher()
	if err != nil {
		return fmt.Errorf("create watcher: %w", err)
	}
	defer func() {
		_ = fswatcher.Close()
	}()

	for _, root := range w.config.Roots {
		if err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !d.IsDir() {
				return nil
			}

			return fswatcher.Add(path)
		}); err != nil {
			return err
		}
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		case event, ok := <-fswatcher.Events:
			if !ok {
				return nil
			}

			log.Printf("watcher: %v %s", event.Op, event.Name)

		case err, ok := <-fswatcher.Errors:
			if !ok {
				return nil
			}

			log.Printf("watcher error: %v", err)
		}
	}
}
