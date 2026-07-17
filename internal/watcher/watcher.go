package watcher

import "time"

type Watcher struct {
	config Config

	debounce chan struct{}
}

func New(config Config) *Watcher {
	if config.Debounce <= 0 {
		config.Debounce = 500 * time.Millisecond
	}

	return &Watcher{
		config:   config,
		debounce: make(chan struct{}, 1),
	}
}
