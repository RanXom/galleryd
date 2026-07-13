package api

import "net/http"

func New(config Config) *Server {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}

	srv := &Server{
		config: config,
		mux:    mux,
		http:   server,
	}

	srv.mux.HandleFunc(
		"GET /health",
		srv.handleHealth,
	)

	return srv
}
