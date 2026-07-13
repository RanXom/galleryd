package api

import (
	"github.com/RanXom/galleryd/internal/service"
	"github.com/RanXom/galleryd/internal/thumbnail"
)

// Config configures the HTTP server
type Config struct {
	Address string

	Gallery    service.GalleryService
	Thumbnails *thumbnail.Generator
}
