package user

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	domainSession "github.com/go-park-mail-ru/2023_1_ContentDealers/session/pkg/domain"
	domainUser "github.com/go-park-mail-ru/2023_1_ContentDealers/user/pkg/domain"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	userCreate := userDTO{}
	err := decoder.Decode(&userCreate)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("failed to parse json string from the body: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	if userCreate.AvatarURL == "" {
		userCreate.AvatarURL = "media/avatars/default_avatar.jpg"
	}
	user := domainUser.User{
		Email:        userCreate.Email,
		PasswordHash: userCreate.Password,
		AvatarURL:    userCreate.AvatarURL,
	}

	_, err = h.userGateway.Register(r.Context(), user)
	if err != nil {
		switch {
		case errors.Is(err, domainUser.ErrUserAlreadyExists):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 1, "message":"user already exists"}`)
		case errors.Is(err, domainUser.ErrNotValidEmail):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 2, "message":"email not validated"}`)
		case errors.Is(err, domainUser.ErrNotValidPassword):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 3, "message":"password not validated"}`)

		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	credentials := userDTO{}
	err := decoder.Decode(&credentials)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("failed to parse json string from the body: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	user := domainUser.User{
		Email:        credentials.Email,
		PasswordHash: credentials.Password, // no hash
	}

	// TODO: user, session, err := h.userUseCase.Auth(r.Context(), user)

	user, err = h.userGateway.Auth(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status": 4, "message":"auth wrong credentials"}`)
		return
	}

	session, err := h.sessionGateway.Create(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionCookie := http.Cookie{
		Name:     "session_id",
		Value:    session.ID,
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Path:     "/",
		// FIXME: Secure: true,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domainSession.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domainSession.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.sessionGateway.Delete(r.Context(), session.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	sessionCookie := http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now().AddDate(0, 0, -1),
		HttpOnly: true,
		Path:     "/",
		// FIXME: Secure: true,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) HasAccessContent(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := r.Context()
	originalURI := r.Header.Get("X-Original-URI")
	if originalURI == "" {
		h.logger.WithRequestID(ctx).Trace("not found 'X-Original-URI' for determine access to content")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var cookieString string
	cookie, err := r.Cookie("session_id")
	if err == nil {
		cookieString = cookie.Value
	}
	err = h.userUsecase.HasAccessContent(ctx, originalURI, cookieString)
	if err != nil {
		h.logger.WithRequestID(ctx).Trace(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
