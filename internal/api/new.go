package api

import "net/http"

func New(config Config) *Server {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}

	return &Server{
		config: config,
		http:   server,
	}
}
