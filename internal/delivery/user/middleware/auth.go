package middleware

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/delivery/user"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/pkg/logging"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func NewAuth(sessionUseCase user.SessionUseCase, logger logging.Logger) Auth {
	return Auth{sessionUseCase: sessionUseCase, logger: logger}
}

type Auth struct {
	sessionUseCase user.SessionUseCase
	logger         logging.Logger
}

func (mw *Auth) RequireUnAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionIDRaw, err := r.Cookie("session_id")
		if err != nil {
			mw.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Trace(err)
			handler.ServeHTTP(w, r)
			return
		}

		sessionID, err := uuid.Parse(sessionIDRaw.Value)
		if err != nil {
			mw.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Tracef("failed to parse uuid from the cookie: %w", err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "failed to parse uuid from the cookie"}`)
			return
		}

		session, err := mw.sessionUseCase.Get(r.Context(), sessionID)
		if err == nil && session.ExpiresAt.After(time.Now()) {
			mw.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Trace("user is already logged in")
			w.WriteHeader(http.StatusBadRequest)
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
			w.WriteHeader(http.StatusBadRequest)
			mw.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Trace(err)
			io.WriteString(w, `{"message": "user is not authorized"}`)
			return
		}

		sessionID, err := uuid.Parse(sessionIDRaw.Value)
		if err != nil {
			mw.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Trace(err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "failed to parse uuid from the cookie"}`)
			return
		}

		session, err := mw.sessionUseCase.Get(r.Context(), sessionID)
		if err != nil || session.ExpiresAt.Before(time.Now()) {
			mw.logger.WithFields(logrus.Fields{
				"request_id": r.Context().Value("requestID").(string),
			}).Trace(err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"message": "user session expired"}`)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "session", session)

		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
