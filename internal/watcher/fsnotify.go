package watcher

import (
	"context"
	"fmt"
	"log"
	"os"

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
		if err := watchTree(fswatcher, root); err != nil {
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

			if !event.Has(fsnotify.Create) {
				continue
			}

			info, err := os.Stat(event.Name)
			if err != nil {
				continue
			}

			if !info.IsDir() {
				continue
			}

			if err := watchTree(fswatcher, event.Name); err != nil {
				log.Printf("watch new directory: %v", err)
			}

		case err, ok := <-fswatcher.Errors:
			if !ok {
				return nil
			}

			log.Printf("watcher error: %v", err)
		}
	}
}
