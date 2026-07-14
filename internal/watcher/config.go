package watcher

import "time"

// Config configures a filesystem watcher
type Config struct {
	Roots []string

	Debounce time.Duration
	OnChange func()
}
