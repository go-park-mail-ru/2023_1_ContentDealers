package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
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

func (mv *GeneralMiddleware) AccessLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("access log middleware")
		start := time.Now()
		next.ServeHTTP(w, r)
		mv.logger.WithFields(logrus.Fields{
			"method": r.Method,
			// код ответа...
			"remote_addr": r.RemoteAddr,
			"url":         r.URL.Path,
			"time":        fmt.Sprintf("%d mcs", time.Since(start).Microseconds()),
			"user_agent":  r.UserAgent(),
		}).Info("New request")
	})
}
