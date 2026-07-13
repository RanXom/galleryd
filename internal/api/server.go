package api

import (
	"net/http"

	"github.com/RanXom/galleryd/internal/service"
	"github.com/RanXom/galleryd/internal/thumbnail"
)

// Server exposes the gallery over HTTP
type Server struct {
	config Config

	mux  *http.ServeMux
	http *http.Server

	gallery    service.GalleryService
	thumbnails *thumbnail.Generator
}
