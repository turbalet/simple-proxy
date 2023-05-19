package service

import "net/http"

var hopHeaders = []string{
	"Connection",
	"Keep-Alive",
	"Upgrade",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailers",
	"Transfer-Encoding",
}

func removeHopHeaders(header http.Header) {
	for _, h := range hopHeaders {
		header.Del(h)
	}
}
