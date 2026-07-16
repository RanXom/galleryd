package watcher

import (
	"context"
	"time"
)

// notify schedules a debounced notification.
//
// Multiple filesystem events collapse into a single callback.
func (w *Watcher) notify() {
	if w.config.OnChange == nil {
		return
	}

	select {
	case w.debounce <- struct{}{}:
	default:
	}
}

func (w *Watcher) debounceLoop(ctx context.Context) {
	if w.config.OnChange == nil {
		return
	}

	timer := time.NewTimer(time.Hour)

	if !timer.Stop() {
		select {
		case <-timer.C:
		default:
		}
	}

	var pending bool

	for {
		select {
		case <-ctx.Done():
			return

		case <-w.debounce:
			pending = true

			if !timer.Stop() {
				select {
				case <-timer.C:
				default:
				}
			}

			timer.Reset(w.config.Debounce)

		case <-timer.C:
			if pending {
				pending = false
				w.config.OnChange()
			}
		}
	}
}
