package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *handler) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/jobs/{id}", h.handleGetJob)
	router.Post("/jobs", h.handlePostJob)

	return router
}
