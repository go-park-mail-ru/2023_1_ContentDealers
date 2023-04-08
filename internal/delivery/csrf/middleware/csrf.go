package middleware

import (
	"io"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
)

const headerCSRF = "csrf-token"

type CSRF struct {
	csrfUseCase csrf.CSRF
}

func NewCSRF(csrfUseCase csrf.CSRF) CSRF {
	return CSRF{csrfUseCase: csrfUseCase}
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
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "csrf token was not given in header 'csrf-token'"}`)
			return
		}
		sessionRaw := r.Context().Value("session")
		session, ok := sessionRaw.(domain.Session)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		isValid, err := mc.csrfUseCase.Check(session, CSRFToken)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if !isValid {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "csrf token is invalid"}`)
			return
		}
		log.Println("csrf token is valid")
		handler.ServeHTTP(w, r)
	})
}
