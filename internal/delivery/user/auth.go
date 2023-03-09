package user

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	credentials := domain.UserCredentials{}
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}
	newUser, err := h.userUseCase.Register(credentials)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrUserAlreadyExists):
			w.WriteHeader(http.StatusConflict)
			io.WriteString(w, `{"status":409}`)
		default:
			// TODO: данные не прошли валидацию, нужен кастомный тип ошибки
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{"status":400}`)
		}
		return
	}

	log.Printf("New User - %v", newUser)
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, `{"status":201}`)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	credentials := domain.UserCredentials{}
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status":400}`)
		return
	}

	user, err := h.userUseCase.Auth(credentials)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"status":404}`)
		return
	}

	session, err := h.sessionUseCase.Create(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
	}

	sessionCookie := http.Cookie{
		Name:     "session_id",
		Value:    session.ID.String(),
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		Path:     "/",
		// FIXME: Secure: true,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status":200}`)
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
		return
	}

	err := h.sessionUseCase.Delete(session.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
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
	io.WriteString(w, `{"status":200}`)
}
