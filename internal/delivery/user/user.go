package user

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/contract"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
)

type Handler struct {
	userUseCase    contract.UserUseCase
	sessionUseCase contract.SessionUseCase
}

func NewHandler(user contract.UserUseCase, session contract.SessionUseCase) Handler {
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
		io.WriteString(w, `{"status":500}`)
		return
	}

	user, err := h.userUseCase.GetByID(session.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
	}

	response, err := json.Marshal(map[string]interface{}{
		"status": http.StatusOK,
		"body": map[string]interface{}{
			"user": map[string]string{
				"email": user.Email,
			},
		},
	})

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, `{"status":500}`)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
