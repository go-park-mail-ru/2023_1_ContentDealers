package middleware

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (mv *GeneralMiddleware) Panic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				mv.logger.WithFields(logrus.Fields{
					"method":     r.Method,
					"url":        r.URL.Path,
					"request_id": r.Context().Value("requestID").(string),
				}).Panicf("recovered %s", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
