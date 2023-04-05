package middleware

import (
	"io"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user/csrf"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

const headerCSRF = "csrf-token"

type CSRF struct {
	cryptToken csrf.CryptToken
}

func NewCSRF(cryptToken csrf.CryptToken) CSRF {
	return CSRF{cryptToken: cryptToken}
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
		isValid, err := mc.cryptToken.Check(session, CSRFToken)
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
