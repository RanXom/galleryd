package api

import (
	"net/http"
)

// Server exposes the gallery over HTTP
type Server struct {
	config Config

	mux  *http.ServeMux
	http *http.Server
}
