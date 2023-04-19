package middleware

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/content/internal/repository/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

const headerCSRF = "csrf-token"

type CSRF struct {
	csrfUseCase csrf.CSRF
	logger      logging.Logger
}

func NewCSRF(csrfUseCase csrf.CSRF, logger logging.Logger) CSRF {
	return CSRF{csrfUseCase: csrfUseCase, logger: logger}
}

// TODO: RequireCSRF должен обрабатывать запрос после RequireAuth
// RequireCSRF обрабатывает только POST запросы
func (mc *CSRF) RequireCSRF(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			handler.ServeHTTP(w, r)
			return
		}
		CSRFToken := r.Header.Get(headerCSRF)
		if CSRFToken == "" {
			msg := "csrf token was not given in header 'csrf-token'"
			mc.logger.Trace(msg)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, fmt.Sprintf(`{"message": "%s"}`, msg))
			return
		}
		sessionRaw := r.Context().Value("session")
		session, ok := sessionRaw.(domain.Session)
		if !ok {
			mc.logger.Trace(domain.ErrSessionInvalid)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		isValid, err := mc.csrfUseCase.Check(session, CSRFToken)
		if err != nil || !isValid {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "csrf token is invalid"}`)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
