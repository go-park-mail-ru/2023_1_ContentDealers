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

type ResponseWriter_CaptureStatus struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter_CaptureStatus(w http.ResponseWriter) *ResponseWriter_CaptureStatus {
	return &ResponseWriter_CaptureStatus{w, 0}
}

func (rw *ResponseWriter_CaptureStatus) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (mv *GeneralMiddleware) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := mv.generateRequestID()
		mv.logger.WithFields(logrus.Fields{
			"method":      r.Method,
			"origin":      r.Header.Get("Origin"),
			"remote_addr": r.Header.Get("X-Real-IP"),
			"url":         r.URL.Path,
			"user_agent":  r.UserAgent(),
			"request_id":  requestID,
		}).Debug("accepted_by_api_gateway")
		ctx := context.WithValue(r.Context(), "requestID", requestID)
		start := time.Now()

		rw := NewResponseWriter_CaptureStatus(w)
		next.ServeHTTP(rw, r.WithContext(ctx))
		mv.logger.WithFields(logrus.Fields{
			"status_http": rw.statusCode,
			"request_id":  requestID,
			"time":        fmt.Sprintf("%d mcs", time.Since(start).Microseconds()),
		}).Debug("released_by_api_gateway")
	})
}
