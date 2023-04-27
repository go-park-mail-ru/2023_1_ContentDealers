package user

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	userCreate := userCreateDTO{}
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
	user := domain.User{
		Email:        userCreate.Email,
		PasswordHash: userCreate.Password,
		AvatarURL:    userCreate.AvatarURL,
	}

	_, err = h.userGateway.Register(r.Context(), user)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserAlreadyExists):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 1, "message":"user already exists"}`)
		case errors.Is(err, domain.ErrNotValidEmail):
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status": 2, "message":"email not validated"}`)
		case errors.Is(err, domain.ErrNotValidPassword):
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
	credentials := userCredentialsDTO{}
	err := decoder.Decode(&credentials)
	if err != nil {
		h.logger.WithRequestID(ctx).Tracef("failed to parse json string from the body: %w", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"message":"failed to parse json string from the body"}`)
		return
	}

	user := domain.User{
		Email:        credentials.Email,
		PasswordHash: credentials.Password, // no hash
	}

	user, err = h.userGateway.Auth(r.Context(), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status": 4, "message":"auth wrong credentials"}`)
		return
	}

	session, err := h.sessionUseCase.Create(r.Context(), user)
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
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		h.logger.WithRequestID(ctx).Trace(domain.ErrSessionInvalid)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err := h.sessionUseCase.Delete(r.Context(), session.ID)
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
