package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) handlePhotos(
	w http.ResponseWriter,
	r *http.Request,
) {
	photos, err := s.gallery.Gallery(r.Context())
	if err != nil {
		log.Printf("build gallery: %v", err)

		writeError(
			w,
			http.StatusInternalServerError,
			"failed to build gallery",
		)
		return
	}

	response := make([]photoResponse, 0, len(photos))

	for _, photo := range photos {
		response = append(response, photoResponse{
			ID:           photo.ID,
			DateTaken:    photo.DateTaken,
			Width:        photo.Width,
			Height:       photo.Height,
			ThumbnailURL: thumbnailURL(photo.ID),
			PhotoURL:     photoURL(photo.ID),
		})
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("encode photos response: %v", err)

		writeError(
			w,
			http.StatusInternalServerError,
			"failed to encode response",
		)
	}
}
