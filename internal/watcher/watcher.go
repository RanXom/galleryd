package watcher

import "context"

type Watcher struct {
	config Config
}

func New(config Config) *Watcher {
	return &Watcher{
		config: config,
	}
}

// Run starts the watcher.
//
// Placeholder implementation.
func (w *Watcher) Run(ctx context.Context) error {
	<-ctx.Done()
	return nil
}
