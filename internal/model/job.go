package model

import "net/http"

type Job struct {
	ID            string      `json:"id"`
	Status        int         `json:"status"`
	URL           string      `json:"url"`
	ContentLength int64       `json:"content_length"`
	ReqMethod     string      `json:"method"`
	ReqHeaders    http.Header `json:"req_headers"`
	RespHeaders   http.Header `json:"resp_headers"`
}
