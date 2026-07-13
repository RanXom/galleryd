package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RanXom/galleryd/internal/api"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	server := api.New(api.Config{
		Address: ":8082",
	})

	log.Println("galleryd listening on :8082")

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
