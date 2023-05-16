package middleware

import (
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/metrics"
)

func replaceID(path string) string {
	re := regexp.MustCompile(`/(\d+)`)

	matches := re.FindAllStringSubmatch(path, -1)

	for _, match := range matches {
		_, err := strconv.Atoi(match[1])
		if err == nil {
			path = regexp.MustCompile(match[0]).ReplaceAllString(path, "/id")
		}
	}

	return path
}

func (mv *GeneralMiddleware) Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		elapsedSeocnds := time.Since(start).Seconds()

		rw, ok := w.(*ResponseWriter_CaptureStatus)
		if !ok {
			return
		}

		// films/1 -> films/id
		path := replaceID(r.URL.Path)

		method := r.Method
		status := strconv.Itoa(rw.statusCode)

		metrics.HttpRequestsTotal.WithLabelValues("api_gateway", path, method, status).Inc()
		metrics.HttpRequestsDurationHistorgram.WithLabelValues("api_gateway", path, method).Observe(elapsedSeocnds)
	})
}
