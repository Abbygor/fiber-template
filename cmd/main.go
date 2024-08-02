package main

import (
	"context"
	"fiber-template/cmd/httpserver"
	"fiber-template/internal/container"
	"os"

	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()

	log := zerolog.New(os.Stdout).With().Timestamp().Logger()

	dependencies := container.Build(ctx, log)

	server := httpserver.NewServer(&dependencies, &log)
	server.SetErrorHandler()
	server.Routes()

	server.Start()
}
