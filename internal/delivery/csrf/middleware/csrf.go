package middleware

import (
	"fmt"
	"io"
	"net/http"

	csrfDelivery "github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/sirupsen/logrus"
)

type CSRF struct {
	csrfUseCase csrf.CSRF
	logger      logging.Logger
	header      string
}

func NewCSRF(csrfUseCase csrf.CSRF, logger logging.Logger, cfg csrfDelivery.CSRFConfig) CSRF {
	return CSRF{
		csrfUseCase: csrfUseCase,
		logger:      logger,
		header:      cfg.Header,
	}
}

// TODO: RequireCSRF должен обрабатывать запрос после RequireAuth
// RequireCSRF обрабатывает только POST запросы
func (mc *CSRF) RequireCSRF(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			handler.ServeHTTP(w, r)
			return
		}
		CSRFToken := r.Header.Get(mc.header)
		if CSRFToken == "" {
			msg := "csrf token was not given in header 'csrf-token'"
			mc.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Trace(msg)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, fmt.Sprintf(`{"message": "%s"}`, msg))
			return
		}
		sessionRaw := r.Context().Value("session")
		session, ok := sessionRaw.(domain.Session)
		if !ok {
			mc.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Trace(domain.ErrSessionInvalid)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		isValid, err := mc.csrfUseCase.Check(r.Context(), session, CSRFToken)
		if err != nil || !isValid {
			mc.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Tracef("csrf token is invalid: %w, isValid: %d", err, isValid)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "csrf token is invalid"}`)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
