package dto

import "net/http"

type JobResponse struct {
	ID      string      `json:"id"`
	Headers http.Header `json:"headers"`
	Length  int64       `json:"length"`
	Status  int         `json:"status"`
}
