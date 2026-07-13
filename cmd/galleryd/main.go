package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/RanXom/galleryd/internal/api"
	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
	"github.com/RanXom/galleryd/internal/service"
	"github.com/RanXom/galleryd/internal/thumbnail"
)

func main() {
	addr := flag.String("addr", ":8082", "HTTP listen address")

	flag.Parse()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	scanner := scanner.New(scanner.Config{
		Roots: []string{
			filepath.Join(os.Getenv("HOME"), "Pictures"),
		},
	})

	reader := metadata.New()

	builder := gallery.New(reader)

	galleryService := service.New(
		scanner,
		builder,
	)
	if err := galleryService.Load(ctx); err != nil {
		log.Fatal(err)
	}

	thumbnailGenerator, err := thumbnail.New(".cache/galleryd")
	if err != nil {
		log.Fatal(err)
	}

	server := api.New(api.Config{
		Address:    *addr,
		Gallery:    galleryService,
		Thumbnails: thumbnailGenerator,
	})

	log.Printf("galleryd listening on %s", *addr)

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
