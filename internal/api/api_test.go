package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/service"
	"github.com/RanXom/galleryd/internal/thumbnail"
)

type fakeGalleryService struct {
	photos []gallery.Photo
	err    error
}

type fakeThumbnailGenerator struct {
	path string
	err  error
}

func (f fakeThumbnailGenerator) Generate(
	ctx context.Context,
	photo gallery.Photo,
) (thumbnail.Thumbnail, error) {
	if f.err != nil {
		return thumbnail.Thumbnail{}, f.err
	}

	return thumbnail.Thumbnail{
		Path: f.path,
	}, nil
}

func (f fakeGalleryService) Reload(ctx context.Context) error {
	return nil
}

func (f fakeGalleryService) Gallery(ctx context.Context) ([]gallery.Photo, error) {
	return f.photos, f.err
}

func (f fakeGalleryService) Photo(
	ctx context.Context,
	id string,
) (gallery.Photo, error) {
	for _, photo := range f.photos {
		if photo.ID == id {
			return photo, nil
		}
	}

	return gallery.Photo{}, service.ErrPhotoNotFound
}

func TestPhotos(t *testing.T) {
	srv := New(Config{
		Address: ":0",

		Gallery: fakeGalleryService{
			photos: []gallery.Photo{
				{
					ID: "abc123",

					Metadata: metadata.Metadata{
						Width:  100,
						Height: 50,
						DateTaken: time.Date(
							2025,
							1,
							1,
							12,
							0,
							0,
							0,
							time.UTC,
						),
					},
				},
			},
		},
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/photos",
		nil,
	)

	rec := httptest.NewRecorder()

	srv.mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	var response []photoResponse

	if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
		t.Fatalf("decode response: %v", err)
	}

	if len(response) != 1 {
		t.Fatalf(
			"expected 1 photo, got %d",
			len(response),
		)
	}

	photo := response[0]

	if photo.ID != "abc123" {
		t.Fatalf(
			"expected id %q, got %q",
			"abc123",
			photo.ID,
		)
	}

	if photo.ThumbnailURL != "/thumb/abc123" {
		t.Fatalf(
			"unexpected thumbnail url: %q",
			photo.ThumbnailURL,
		)
	}

	if photo.PhotoURL != "/photo/abc123" {
		t.Fatalf(
			"unexpected photo url: %q",
			photo.PhotoURL,
		)
	}
}

func TestThumbnail(t *testing.T) {
	srv := New(Config{
		Address: ":0",

		Gallery: fakeGalleryService{
			photos: []gallery.Photo{
				{
					ID: "abc123",
				},
			},
		},

		Thumbnails: fakeThumbnailGenerator{
			path: "../scanner/testdata/gallery/vac/image.webp",
		},
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/thumb/abc123",
		nil,
	)

	rec := httptest.NewRecorder()

	srv.mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf(
			"expected %d, got %d",
			http.StatusOK,
			rec.Code,
		)
	}

	if got := rec.Header().Get("Content-Type"); got != "image/webp" {
		t.Fatalf(
			"expected image/webp, got %q",
			got,
		)
	}
}
