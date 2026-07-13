package api

import (
	"context"
	"errors"
	"net/http"
)

// Run starts the HTTP server and blocks until it exits.
func (s *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()

		_ = s.http.Shutdown(context.Background())
	}()

	err := s.http.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}
