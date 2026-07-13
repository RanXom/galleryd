package api

import "net/http"

func New(config Config) *Server {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}

	srv := &Server{
		config:  config,
		mux:     mux,
		http:    server,
		gallery: config.Gallery,
	}

	srv.mux.HandleFunc(
		"GET /health",
		srv.handleHealth,
	)

	srv.mux.HandleFunc(
		"GET /api/photos",
		srv.handlePhotos,
	)

	return srv
}
