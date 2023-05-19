package handler

import (
	"net/http"
	"simple-proxy/internal/service"
	"simple-proxy/pkg/logger"
)

type Handler interface {
	RegisterRoutes() http.Handler
}

type handler struct {
	logger  logger.Logger
	service service.Service
}

func NewHandler(logger logger.Logger, service service.Service) Handler {
	return &handler{
		logger:  logger,
		service: service,
	}
}
