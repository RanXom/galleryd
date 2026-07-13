package api

import (
	"errors"
	"log"
	"net/http"

	"github.com/RanXom/galleryd/internal/service"
)

func (s *Server) handleThumbnail(
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

	thumb, err := s.thumbnails.Generate(
		r.Context(),
		photo,
	)
	if err != nil {
		log.Printf("generate thumbnail: %v", err)

		writeError(
			w,
			http.StatusInternalServerError,
			"failed to generate thumbnail",
		)
		return
	}

	http.ServeFile(
		w,
		r,
		thumb.Path,
	)
}
