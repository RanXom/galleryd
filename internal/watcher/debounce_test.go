package watcher

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestDebounceCollapsesNotifications(t *testing.T) {
	done := make(chan struct{}, 1)
	var calls atomic.Int32

	w := New(Config{
		Debounce: 25 * time.Millisecond,
		OnChange: func() {
			calls.Add(1)

			select {
			case done <- struct{}{}:
			default:
			}
		},
	})

	go w.debounceLoop(t.Context())

	for range 10 {
		w.notify()
	}

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for callback")
	}

	if got := calls.Load(); got != 1 {
		t.Fatalf("expected 1 callback, got %d", got)
	}
}

func TestDebounceSeparateBursts(t *testing.T) {
	done := make(chan struct{}, 2)
	var calls atomic.Int32

	w := New(Config{
		Debounce: 20 * time.Millisecond,
		OnChange: func() {
			calls.Add(1)

			select {
			case done <- struct{}{}:
			default:
			}
		},
	})

	go w.debounceLoop(t.Context())

	w.notify()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for first callback")
	}

	w.notify()

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for second callback")
	}

	if got := calls.Load(); got != 2 {
		t.Fatalf("expected 2 callbacks, got %d", got)
	}
}
