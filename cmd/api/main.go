package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"simple-proxy/internal/config"
	"simple-proxy/internal/handler"
	"simple-proxy/internal/repository"
	"simple-proxy/internal/server"
	"simple-proxy/internal/service"
	"simple-proxy/pkg/logger"
	"syscall"
)

func main() {
	var cfg config.Config
	flag.IntVar(&cfg.Port, "Port", 8000, "Proxy server port")

	logger, err := logger.RegisterLog("dev", "dev")
	if err != nil {
		log.Fatalf("err initializing logger: %v", err)
	}

	repository := repository.NewRepository()
	service := service.NewService(logger, repository)
	handler := handler.NewHandler(logger, service)
	srv := server.NewServer(cfg.Port, handler)
	go func() {
		if err := srv.Run(); err != nil {
			logger.Fatalf("err starting server: %v", err)
		}
	}()

	logger.Infof("listening :%d", cfg.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info(context.Background(), "server shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
