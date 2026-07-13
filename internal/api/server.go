package api

import "net/http"

// Server exposes the gallery over HTTP
type Server struct {
	config Config

	http *http.Server
}
