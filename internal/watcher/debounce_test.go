package watcher

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestDebounceCollapsesNotifications(t *testing.T) {
	var calls atomic.Int32

	w := New(Config{
		Debounce: 25 * time.Millisecond,
		OnChange: func() {
			calls.Add(1)
		},
	})

	go w.debounceLoop(t.Context())

	for range 10 {
		w.notify()
	}

	time.Sleep(75 * time.Millisecond)

	if got := calls.Load(); got != 1 {
		t.Fatalf("expected 1 callback, got %d", got)
	}
}

func TestDebounceSeparateBursts(t *testing.T) {
	var calls atomic.Int32

	w := New(Config{
		Debounce: 20 * time.Millisecond,
		OnChange: func() {
			calls.Add(1)
		},
	})

	go w.debounceLoop(t.Context())

	w.notify()

	time.Sleep(40 * time.Millisecond)

	w.notify()

	time.Sleep(40 * time.Millisecond)

	if got := calls.Load(); got != 2 {
		t.Fatalf("expected 2 callbacks, got %d", got)
	}
}
