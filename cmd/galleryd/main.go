package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RanXom/galleryd/internal/api"
	"github.com/RanXom/galleryd/internal/gallery"
	"github.com/RanXom/galleryd/internal/metadata"
	"github.com/RanXom/galleryd/internal/scanner"
	"github.com/RanXom/galleryd/internal/service"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	scanner := scanner.New(scanner.Config{
		Roots: []string{
			".",
		},
	})

	reader := metadata.New()

	builder := gallery.New(reader)

	gallerySerice := service.New(
		scanner,
		builder,
	)

	server := api.New(api.Config{
		Address: ":8082",
		Gallery: gallerySerice,
	})

	log.Println("galleryd listening on :8082")

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
