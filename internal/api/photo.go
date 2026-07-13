package api

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/RanXom/galleryd/internal/service"
)

type photoResponse struct {
	ID           string    `json:"id"`
	DateTaken    time.Time `json:"dateTaken"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	ThumbnailURL string    `json:"thumbnail"`
	PhotoURL     string    `json:"photo"`
}

func (s *Server) handlePhoto(
	w http.ResponseWriter,
	r *http.Request,
) {
	id := r.PathValue("id")

	photo, err := s.gallery.Photo(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrPhotoNotFound) {
			writeError(
				w,
				http.StatusNotFound,
				"photo not found",
			)
			return
		}

		log.Printf("lookup photo: %v", err)

		writeError(
			w,
			http.StatusInternalServerError,
			"failed to lookup photo",
		)
		return
	}

	http.ServeFile(
		w,
		r,
		photo.Path,
	)
}
