package server

import (
	"context"
	"fmt"
	"net/http"
	"simple-proxy/internal/handler"
	"time"
)

type ProxyServer struct {
	server *http.Server
}

func NewServer(port int, handler handler.Handler) *ProxyServer {
	return &ProxyServer{
		server: &http.Server{
			Addr:         fmt.Sprintf(":%d", port),
			Handler:      handler.RegisterRoutes(),
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}}
}

func (ps *ProxyServer) Run() error {
	return ps.server.ListenAndServe()
}

func (ps *ProxyServer) Shutdown(ctx context.Context) error {
	return ps.server.Shutdown(ctx)
}
