package handler

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func responseJSON(w http.ResponseWriter, statusCode int, resp any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(resp)
}

func isValidURL(inUrl string) bool {
	_, err := url.ParseRequestURI(inUrl)
	if err != nil {
		return false
	}
	return true
}
