package middleware

import (
	"net/http"
)

type sampleData struct {
	Data string `json:"data"`
}

// Middleware type is an adapter to allow the use of ordinary functions as
// handlers.
type Middleware func(http.Handler) http.Handler

// HTTPHandlerMock structure is usead for injecting dependency of type
// http.Handler.
type HTTPHandlerMock struct {
	w http.ResponseWriter
	r *http.Request
}

func (h *HTTPHandlerMock) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
