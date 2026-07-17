package watcher

import (
	"context"
	"testing"
	"time"
)

func TestRunReturnsOnContextCancellation(t *testing.T) {
	w := New(Config{})

	ctx, cancel := context.WithCancel(context.Background())

	done := make(chan error, 1)

	go func() {
		done <- w.Run(ctx)
	}()

	cancel()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("expected nil, got %v", err)
		}

	case <-time.After(time.Second):
		t.Fatal("watcher did not exit")
	}
}
