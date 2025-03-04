package server

import (
	"context"
	"fmt"
	"log"
	"net"

	test "example.com/m/pkg/api/order/api"
	"example.com/m/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Config struct {
	Host string
	Port string
}

type Router struct {
	Config  Config
	Handler *Handler
	Server  *grpc.Server
	Lis     *net.Listener
}

func NewRouter(cfg Config, h Handler) *Router {
	server := grpc.NewServer()
	test.RegisterOrderServiceServer(server, h)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	return &Router{
		Config: cfg,
		Server: server,
		Lis:    &lis,
	}
}

func (r *Router) Run(ctx context.Context) {
	if err := r.Server.Serve(*r.Lis); err != nil {
		logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve", zap.Error(err))
	}
	logger.GetLoggerFromCtx(ctx).Info(ctx, "Server is runing")
}
