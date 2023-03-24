package middleware

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/google/uuid"
)

func NewAuth(sessionUseCase user.SessionUseCase) Auth {
	return Auth{sessionUseCase: sessionUseCase}
}

type Auth struct {
	sessionUseCase user.SessionUseCase
}

func (mw *Auth) RequireUnAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIDRaw, err := r.Cookie("session_id")
		if err != nil {
			handler.ServeHTTP(w, r)
			return
		}

		sessionID, err := uuid.Parse(sessionIDRaw.Value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "failed to parse uuid from the cookie"}`)
			return
		}

		session, err := mw.sessionUseCase.Get(sessionID)
		if err == nil && session.ExpiresAt.After(time.Now()) {
			w.WriteHeader(http.StatusForbidden)
			io.WriteString(w, `{"message": "user is already logged in"}`)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (mw *Auth) RequireAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIDRaw, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"message": "user is not authorized"}`)
			return
		}

		sessionID, err := uuid.Parse(sessionIDRaw.Value)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "failed to parse uuid from the cookie"}`)
			return
		}

		session, err := mw.sessionUseCase.Get(sessionID)
		if err != nil || session.ExpiresAt.Before(time.Now()) {
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"message": "user session expired"}`)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "session", session)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
