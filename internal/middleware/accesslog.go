package middleware

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func NewGeneral(logger logging.Logger) GeneralMiddleware {
	return GeneralMiddleware{
		logger: logger,
	}
}

type GeneralMiddleware struct {
	logger logging.Logger
}

func (mv *GeneralMiddleware) generateRequestID() string {
	return uuid.New().String()
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, 0}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (mv *GeneralMiddleware) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := mv.generateRequestID()
		mv.logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"origin":      r.Header.Get("Origin"),
			"remote_addr": r.RemoteAddr,
			"url":         r.URL.Path,
			"user_agent":  r.UserAgent(),
			"request_id":  requestID,
		}).Debug("Start processing request")
		ctx := context.WithValue(r.Context(), "requestID", requestID)
		start := time.Now()

		loggingRW := NewLoggingResponseWriter(w)
		next.ServeHTTP(loggingRW, r.WithContext(ctx))
		mv.logger.WithFields(logrus.Fields{
			"status_http": loggingRW.statusCode,
			"request_id":  requestID,
			"time":        fmt.Sprintf("%d mcs", time.Since(start).Microseconds()),
		}).Debug("End processing request")
	})
}
