package delivery

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/repository"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase"
)

type UserHandler struct {
	userUseCase    *usecase.UserUseCase
	sessionUseCase *usecase.SessionUseCase
}

func NewUserHandler(useCase *usecase.UserUseCase, sessionUseCase *usecase.SessionUseCase) UserHandler {
	return UserHandler{
		userUseCase:    useCase,
		sessionUseCase: sessionUseCase,
	}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	credentials := domain.UserCredentials{}
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status": 400}`)
		return
	}

	newUser, err := h.userUseCase.RegisterUser(credentials)
	if err != nil {
		switch err {
		case repository.ErrUserAlreadyExists:
			w.WriteHeader(http.StatusConflict)
			io.WriteString(w, `{"status": 409}`)
		default:
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{"status": 500}`)
		}
		return
	}

	log.Printf("New User - %v", newUser)
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, `{"status": 201}`)
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	credentials := domain.UserCredentials{}
	err := decoder.Decode(&credentials)
	if err != nil {
		log.Printf("error while unmarshalling JSON: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		io.WriteString(w, `{"status": 400}`)
		return
	}

	user, err := h.userUseCase.AuthUser(credentials)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		io.WriteString(w, `{"status": 404}`)
		return
	}

	session, err := h.sessionUseCase.CreateSession(user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500}`)
	}

	sessionCookie := http.Cookie{
		Name:     "session_id",
		Value:    session.ID.String(),
		Expires:  session.ExpiresAt,
		HttpOnly: true,
		// FIXME: Secure: true,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": 200}`)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500}`)
		return
	}

	err := h.sessionUseCase.DeleteSession(session.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500}`)
		return
	}

	sessionCookie := http.Cookie{
		Name:     "session_id",
		Value:    "",
		Expires:  time.Now().AddDate(0, 0, -1),
		HttpOnly: true,
		// FIXME: Secure: true,
	}

	http.SetCookie(w, &sessionCookie)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"status": 200}`)
}

func (h *UserHandler) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500}`)
		return
	}

	user, err := h.userUseCase.GetById(session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status": 500}`)
	}

	responseBody, err := json.Marshal(map[string]interface{}{
		"status": http.StatusOK,
		"body": map[string]interface{}{
			"user": user,
		},
	})

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, string(responseBody))
}
