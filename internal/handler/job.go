package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"simple-proxy/internal/model/dto"
	"simple-proxy/internal/service"
)

func (h *handler) handleGetJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	job, err := h.service.GetJob(id)
	if err != nil {
		if errors.Is(err, service.ErrJobNotFound) {
			responseJSON(w, http.StatusNotFound, ErrResponse{
				Msg: "job with given id not found",
			})
			return
		}
		responseJSON(w, http.StatusInternalServerError, ErrResponse{
			Msg: "internal server error",
		})
		return
	}

	responseJSON(w, http.StatusOK, job)
}

func (h *handler) handlePostJob(w http.ResponseWriter, r *http.Request) {
	in := dto.JobRequest{}
	err := json.NewDecoder(r.Body).Decode(&in)
	if err != nil {
		responseJSON(w, http.StatusBadRequest, ErrResponse{
			Msg: "invalid body",
		})
		return
	}

	if !isValidURL(in.URL) {
		responseJSON(w, http.StatusBadRequest, ErrResponse{
			Msg: "invalid url",
		})
		return
	}

	jobResponse, err := h.service.CreateJob(in)
	if err != nil {
		responseJSON(w, http.StatusInternalServerError, ErrResponse{
			Msg: "internal server error"})
		return
	}

	responseJSON(w, http.StatusOK, jobResponse)
}
