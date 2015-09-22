package middleware

import (
	"net/http"
)

// SetJSONHeader middleware sets response Content-Type to application/json.
func SetJSONHeader(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
