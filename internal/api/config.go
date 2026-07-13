package api

import (
	"github.com/RanXom/galleryd/internal/service"
)

// Config configures the HTTP server
type Config struct {
	Address string

	Gallery    service.GalleryService
	Thumbnails ThumbnailGenerator
}
