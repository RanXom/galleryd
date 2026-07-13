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
		log.Printf("gallery: %v", err)

		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
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
			ThumbnailURL: "/thumb/" + photo.ID,
			PhotoURL:     "/photo/" + photo.ID,
		})
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(
			w,
			http.StatusText(http.StatusInternalServerError),
			http.StatusInternalServerError,
		)
	}
}
