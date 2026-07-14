package watcher

type Watcher struct {
	config Config
}

func New(config Config) *Watcher {
	return &Watcher{
		config: config,
	}
}
