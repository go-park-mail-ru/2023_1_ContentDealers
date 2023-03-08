package middleware

import (
	"net/http"
	"strings"
)

func SetContentTypeJSON(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		handler.ServeHTTP(w, r)
	})
}

func ValidateRequestContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType == "" || strings.Contains(contentType, "application/json") {
			handler.ServeHTTP(w, r)
			return
		}
		http.Error(w, `{"status":400}`, http.StatusBadRequest)
	})
}
