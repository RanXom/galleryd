package service

import (
	"context"
	"errors"

	"github.com/RanXom/galleryd/internal/gallery"
)

var ErrPhotoNotFound = errors.New("photo not found")

// Photo returns a single photo by its ID.
func (s *galleryService) Photo(
	ctx context.Context,
	id string,
) (gallery.Photo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	select {
	case <-ctx.Done():
		return gallery.Photo{}, ctx.Err()
	default:
	}

	photo, ok := s.index.byID[id]
	if !ok {
		return gallery.Photo{}, ErrPhotoNotFound
	}

	return photo, nil
}
