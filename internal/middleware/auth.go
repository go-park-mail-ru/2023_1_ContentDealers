package middleware

import (
	"context"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

func NewAuthMiddleware(sessionUseCase *usecase.SessionUseCase) AuthMiddleware {
	return AuthMiddleware{sessionUseCase: sessionUseCase}
}

type AuthMiddleware struct {
	sessionUseCase *usecase.SessionUseCase
}

func (mw *AuthMiddleware) UnAuthorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIDRaw, err := r.Cookie("session_id")
		if err != nil {
			handler.ServeHTTP(w, r)
			return
		}

		sessionID, err := uuid.Parse(sessionIDRaw.Value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 400}`)
			return
		}

		session, err := mw.sessionUseCase.GetSession(sessionID)
		if err == nil && session.ExpiresAt.After(time.Now()) {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, `{"status": 403}`)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (mw *AuthMiddleware) Authorized(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIDRaw, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"status": 401}`)
			return
		}

		sessionID, err := uuid.Parse(sessionIDRaw.Value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 400}`)
			return
		}

		session, err := mw.sessionUseCase.GetSession(sessionID)
		if err != nil || session.ExpiresAt.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"status": 401}`)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "session", session)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
