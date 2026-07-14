package api

import (
	"log"
	"net/http"
)

func (s *Server) handleReload(
	w http.ResponseWriter,
	r *http.Request,
) {
	if err := s.gallery.Reload(r.Context()); err != nil {
		log.Printf("reload gallery: %v", err)

		writeError(
			w,
			http.StatusInternalServerError,
			"failed to reload gallery",
		)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}
