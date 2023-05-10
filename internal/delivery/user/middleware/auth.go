package middleware

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
)

func NewAuth(sessionGateway user.SessionGateway, logger logging.Logger) Auth {
	return Auth{sessionGateway: sessionGateway, logger: logger}
}

type Auth struct {
	sessionGateway user.SessionGateway
	logger         logging.Logger
}

func (mw *Auth) RequireUnAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sessionIDRaw, err := r.Cookie("session_id")
		if err != nil {
			handler.ServeHTTP(w, r)
			return
		}

		sessionID := sessionIDRaw.Value

		session, err := mw.sessionGateway.Get(r.Context(), sessionID)
		if err == nil && session.ExpiresAt.After(time.Now()) {
			mw.logger.WithRequestID(ctx).Trace("user is already logged in")
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "user is already logged in"}`)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func (mw *Auth) RequireAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		sessionIDRaw, err := r.Cookie("session_id")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			mw.logger.WithRequestID(ctx).Trace(err)
			io.WriteString(w, `{"message": "user is not authorized"}`)
			return
		}

		sessionID := sessionIDRaw.Value

		session, err := mw.sessionGateway.Get(r.Context(), sessionID)

		if err != nil || session.ExpiresAt.Before(time.Now()) {
			mw.logger.WithRequestID(ctx).Trace(err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "user session expired"}`)
			return
		}

		ctx = context.WithValue(ctx, "session", session)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
