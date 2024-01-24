package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Bubotka/grpc-chat/config"
	"github.com/Bubotka/grpc-chat/server"
)

func main() {
	cfg := config.MustConfig()

	grpcServer := server.NewChatGptServer(cfg, "")

	grpcAddr := fmt.Sprintf(":%s", cfg.GRPC.Port)
	grpcListener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer func() { _ = grpcListener.Close() }()

	go func() {
		err := grpcServer.Start(grpcListener)
		if err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	select {
	case v := <-done:
		log.Println("shutting down server with signal:", v)
	case ctxDone := <-ctx.Done():
		log.Println("shutting down server with context:", ctxDone)
	}
}
