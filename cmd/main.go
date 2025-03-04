package main

import (
	"context"

	orderrepository "example.com/m/internal/order/repository"
	orderservice "example.com/m/internal/order/service"
	server "example.com/m/internal/transport/grpc"
	"example.com/m/pkg/logger"
)

func main() {
	ctx := context.Background()

	ctx, _ = logger.NewLogger(ctx)

	repo := orderrepository.NewRepository()

	service := orderservice.NewService(repo)

	handler := server.NewHandler(service)

	cfg := server.Config{
		Host: "localhost",
		Port: "50051",
	}

	r := server.NewRouter(cfg, *handler)

	r.Run(ctx)
}
