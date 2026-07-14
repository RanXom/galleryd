package watcher

import "time"

// Configures a filesystem watcher
type Config struct {
	Roots []string

	Debounce time.Duration
	OnChange func()
}
