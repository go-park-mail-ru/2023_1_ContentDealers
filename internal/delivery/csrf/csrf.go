package csrf

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/domain"
	"github.com/go-park-mail-ru/2023_1_ContentDealers/internal/usecase/csrf"
)

const ExpirationTimeCSRF = 2 * time.Hour

type Handler struct {
	csrfUseCase csrf.CSRF
}

func NewHandler(csrfUseCase csrf.CSRF) Handler {
	return Handler{
		csrfUseCase: csrfUseCase,
	}
}

// @Summary CSRF
// @Tags user
// @Description Получить CSRF токен
// @Description Необходимы куки
// @Produce  json
// @Success 200 {object} tokenDTO
// @Failure 400
// @Failure 500
// @Router /user/csrf [get]
func (h *Handler) GetCSRF(w http.ResponseWriter, r *http.Request) {
	sessionRaw := r.Context().Value("session")
	session, ok := sessionRaw.(domain.Session)
	if !ok {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	token, err := h.csrfUseCase.Create(session, time.Now().Add(ExpirationTimeCSRF).Unix())
	if err != nil {
		// log "csrf token creation error"
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(map[string]interface{}{
		"body": map[string]interface{}{
			"csrf-token": token,
		},
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
