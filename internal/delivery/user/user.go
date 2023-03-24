package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Handler struct {
	userUseCase    UserUseCase
	sessionUseCase SessionUseCase
}

func NewHandler(user UserUseCase, session SessionUseCase) Handler {
	return Handler{
		userUseCase:    user,
		sessionUseCase: session,
	}
}

func (h *Handler) Info(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	ctx := r.Context()
	sessionRaw := ctx.Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	user, err := h.userUseCase.GetByID(session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"user": map[string]string{
				"email": user.Email,
			},
		},
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
