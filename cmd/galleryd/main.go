package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RanXom/galleryd/internal/api"
	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
	"github.com/RanXom/galleryd/internal/service"
	"github.com/RanXom/galleryd/internal/thumbnail"
	"github.com/RanXom/galleryd/internal/watcher"
)

func main() {
	cacheDir := flag.String(
		"cache-dir",
		defaultCacheDir(),
		"thumbnail cache directory",
	)

	addr := flag.String(
		"addr",
		":8082",
		"HTTP listen address (e.g. :8082, localhost:8082, 0.0.0.0:8082)",
	)

	var dirs stringSliceFlag

	flag.Var(
		&dirs,
		"dir",
		"gallery directory (may be specified multiple times)",
	)

	flag.Parse()

	if len(dirs) == 0 {
		dirs = append(
			dirs,
			defaultPicturesDir(),
		)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	scanner := scanner.New(scanner.Config{
		Roots: dirs,
	})

	reader := metadata.New()

	builder := gallery.New(reader)

	galleryService := service.New(
		scanner,
		builder,
	)
	if err := galleryService.Reload(ctx); err != nil {
		log.Fatal(err)
	}

	fswatcher := watcher.New(watcher.Config{
		Roots: dirs,
	})

	thumbnailGenerator, err := thumbnail.New(*cacheDir)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		if err := fswatcher.Run(ctx); err != nil {
			log.Printf("watcher: %v", err)
		}
	}()

	server := api.New(api.Config{
		Address:    *addr,
		Gallery:    galleryService,
		Thumbnails: thumbnailGenerator,
	})

	log.Printf("galleryd listening on %s", listenURL(*addr))

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
